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
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func GetAuthGroup(e *echo.Group) {
	g := e.Group("/auths")
	g.GET("", getAuths, utils.CheckAutorization)
	g.POST("", addAuth, utils.CheckAutorization)
	g.DELETE("/:id", deleteAuthById, utils.CheckAutorization)
	g.GET("/:id", getAuthById, utils.CheckAutorization)
	g.GET("/:provider/login", authWithOAuth)
	g.GET("/:provider/callback", authWithOAuthCallback)
	g.GET("/login", authWithUser)
}

func authWithUser(c echo.Context) error {
	user := new(models.UserLogin)
	if err := utils.BindAndValidateObject(c, user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userData, err := services.LoginWithUser(c.Get("db").(*gorm.DB), *user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, userData)
}

func authWithOAuth(c echo.Context) error {
	oauthConfig := auth.GetAuthConfig(c.Param("provider"))
	if oauthConfig == nil {
		return echo.NewHTTPError(http.StatusBadRequest, c.Param("provider"))
	}
	oauthState := generateStateOauthCookie(c)
	sess := utils.GetSession(c)
	sess.Values["redirect"] = c.QueryParam("redirect_url")
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

	if c.FormValue("state") != oauthState.Value {
		sess.Values["redirect"] = nil
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusUnauthorized, redirect.(string))
	}
	data, err := getUserDataFromProvider(c.FormValue("code"), oauthConfig, c.Param("provider"))
	if err != nil {
		sess.Values["redirect"] = nil
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusUnauthorized, redirect.(string))
	}
	sess.Values["redirect"] = nil
	sess.Values["userinfo"] = data
	log.Println(data)
	sess.Save(c.Request(), c.Response())
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

// type ErrorDetails struct {
// 	Error            string `json:"error"`
// 	ErrorDescription string `json:"error_description"`
// }
// if rErr, ok := err.(*oauth2.RetrieveError); ok {
// 	details := new(ErrorDetails)
// 	if err := json.Unmarshal(rErr.Body, details); err != nil {
// 		panic(err)
// 	}

// 	log.Println(details.Error, details.ErrorDescription)
// }

func getUserDataFromProvider(code string, oauthConfig *oauth2.Config, provider string) (*models.UserData, error) {
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
	return c.JSON(http.StatusOK, auths)
}

func addAuth(c echo.Context) error {
	auth := new(models.Auth)
	if err := utils.BindAndValidateObject(c, auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.AddAuth(c.Get("db").(*gorm.DB), *auth); err != nil {
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	var auth models.Auth
	if err := utils.BindAndValidateObject(c, &auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.ReplaceAuth(db, id, auth); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
