package controllers

import (
	"apprit/store/api/v1/auth"
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/services"
	"apprit/store/api/v1/utils"
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func GetAuthGroup(e *echo.Group) {
	g := e.Group("/auths")
	g.GET("", getAuths, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.POST("", addAuth, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.DELETE("/:id", deleteAuthById, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.GET("/:id", getAuthById, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.GET("/:provider/login", authWithOAuth)
	g.GET("/:provider/callback", authWithOAuthCallback)
	g.POST("/login", authLogin)
	g.POST("/register", authRegister)
	g.GET("/logout", authLogout, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.POST("/token", authTokenLocal)
}

func authTokenLocal(c echo.Context) error {
	user := new(models.UserLogin)
	if err := utils.BindAndValidateObject(c, user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userData, err := services.LoginWithUser(c.Get("db").(*gorm.DB), *user)
	if err != nil {
		return err
	}
	t, err := auth.CreateToken(userData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func authLogout(c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "userinfo", Value: "", HttpOnly: true, SameSite: http.SameSiteStrictMode, Expires: time.Unix(0, 0), Path: "/"})
	return c.NoContent(http.StatusOK)
}

func authLogin(c echo.Context) error {
	user := new(models.UserLogin)
	if err := utils.BindAndValidateObject(c, user); err != nil {
		c.SetCookie(&http.Cookie{Name: "login_state", Value: "failure"})
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userData, err := services.LoginWithUser(c.Get("db").(*gorm.DB), *user)
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "login_state", Value: "failure"})
		return err
	}

	t, err := auth.CreateToken(userData)
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "login_state", Value: "failure"})
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.SetCookie(utils.GetTokenCookie(t))
	c.SetCookie(&http.Cookie{Name: "login_state", Value: "success"})
	return c.NoContent(http.StatusOK)
}

func authRegister(c echo.Context) error {
	user := new(models.User)
	if err := utils.BindAndValidateObject(c, user); err != nil {
		c.SetCookie(&http.Cookie{Name: "login_state", Value: "failure"})
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user.Admin = false
	userData, err := services.AddUser(c.Get("db").(*gorm.DB), *user)
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "login_state", Value: "failure"})
		return err
	}

	t, err := auth.CreateToken(models.ConverUserToUserData(userData))
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "login_state", Value: "failure"})
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.SetCookie(utils.GetTokenCookie(t))
	c.SetCookie(&http.Cookie{Name: "login_state", Value: "success"})
	return c.NoContent(http.StatusCreated)
}

func authWithOAuth(c echo.Context) error {
	oauthConfig := auth.GetAuthConfig(c.Param("provider"))
	if oauthConfig == nil {
		return c.Redirect(http.StatusPermanentRedirect, c.Request().Referer())
	}
	redirectUrl := c.QueryParam("redirect_url")
	if redirectUrl == "" {
		return c.Redirect(http.StatusPermanentRedirect, c.Request().Referer())
	}
	oauthState := generateStateOauthCookie(c)
	referer, _ := url.Parse(c.Request().Referer())
	referer.Path = path.Join(referer.Path, c.QueryParam("redirect_url"))
	sess := utils.GetSession(c)
	sess.Values["redirect"] = referer.String()
	sess.Save(c.Request(), c.Response())
	oauthConfigUrl := oauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, oauthConfigUrl)
}

func authWithOAuthCallback(c echo.Context) error {
	oauthConfig := auth.GetAuthConfig(c.Param("provider"))
	if oauthConfig == nil {
		return echo.NewHTTPError(http.StatusBadRequest, c.Param("provider"))
	}

	oauthState, _ := c.Cookie("oauthstate")
	sess := utils.GetSession(c)

	redirect := sess.Values["redirect"]
	if redirect == nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	redirectUrl, _ := url.Parse(redirect.(string))

	if c.FormValue("state") != oauthState.Value {
		sess.Values["redirect"] = nil
		sess.Save(c.Request(), c.Response())
		c.SetCookie(&http.Cookie{Name: "oauthstate", Value: "", Expires: time.Unix(0, 0)})
		c.SetCookie(&http.Cookie{Name: "login_state", Value: "failure", Path: redirectUrl.Path})
		return c.Redirect(http.StatusPermanentRedirect, redirect.(string))
	}
	data, err := getUserDataFromProvider(c.FormValue("code"), oauthConfig, c.Param("provider"))

	log.Println(err)
	if err != nil {
		sess.Values["redirect"] = nil
		sess.Save(c.Request(), c.Response())
		c.SetCookie(&http.Cookie{Name: "oauthstate", Value: "", Expires: time.Unix(0, 0)})
		c.SetCookie(&http.Cookie{Name: "login_state", Value: "failure", Path: redirectUrl.Path})
		return c.Redirect(http.StatusPermanentRedirect, redirect.(string))
	}
	sess.Values["redirect"] = nil
	sess.Save(c.Request(), c.Response())
	db := c.Get("db").(*gorm.DB)
	user, err := services.GetUserByEmail(db, data.Email)
	if err != nil {
		user, err = services.AddUser(db, models.User{Username: data.Name, Email: data.Email, Admin: false})
	}
	t, err := auth.CreateToken(models.ConverUserToUserData(user))
	if err != nil {
		return c.Redirect(http.StatusPermanentRedirect, err.Error())
	}
	c.SetCookie(utils.GetTokenCookie(t))
	c.SetCookie(&http.Cookie{Name: "oauthstate", Value: "", Expires: time.Unix(0, 0)})
	c.SetCookie(&http.Cookie{Name: "login_state", Value: "success", Path: redirectUrl.Path})
	return c.Redirect(http.StatusPermanentRedirect, redirect.(string))
}

func generateStateOauthCookie(c echo.Context) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	c.SetCookie(&cookie)

	return state
}

func getUserDataFromProvider(code string, oauthConfig *oauth2.Config, provider string) (*models.CallbackUserData, error) {
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userData, err := auth.GetUserData(provider, token)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return userData, nil
}

func getAuths(c echo.Context) error {
	auths, err := services.GetAuths(c.Get("db").(*gorm.DB))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, models.ResponseArrayEntity[models.Auth]{Elements: auths})
}

func addAuth(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	auth := models.Auth{}
	if err := utils.BindAndValidateObject(c, &auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if _, err := services.AddAuth(c.Get("db").(*gorm.DB), auth); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getAuthById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	auth, err := services.GetAuthById(c.Get("db").(*gorm.DB), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, auth)
}

func deleteAuthById(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.DeleteAuthById(c.Get("db").(*gorm.DB), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func replaceAuthById(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	auth := models.Auth{}
	if err := utils.BindAndValidateObject(c, &auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.ReplaceAuth(db, id, auth); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
