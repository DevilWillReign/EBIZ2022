package controllers

import (
	"apprit/store/api/v1/auth"
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/services"
	"context"
	"crypto/rand"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func GetAuthGroup(e *echo.Group) {
	g := e.Group("/auths")
	g.GET("", getAuths)
	g.POST("", addAuth)
	g.DELETE("/:id", deleteAuthById)
	g.GET("/:id", getAuthById)
	g.GET("/:provider/login", authWithOAuth)
	g.GET("/:provider/callback", authWithOAuthCallback)
}

func getSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	return sess
}

func authWithOAuth(c echo.Context) error {
	oauthConfig := auth.GetAuthConfig(c.Param("provider"))
	if oauthConfig == nil {
		return echo.NewHTTPError(http.StatusBadRequest, c.Param("provider"))
	}
	oauthState := generateStateOauthCookie(c)
	sess := getSession(c)
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
	sess := getSession(c)

	redirect := sess.Values["redirect"]
	if redirect == nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if c.FormValue("state") != oauthState.Value {
		sess.Values["redirect"] = nil
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusUnauthorized, redirect.(string))
	}
	data, err := getUserDataFromGoogle(c.FormValue("code"), oauthConfig, c.Param("provider"))
	if err != nil {
		sess.Values["redirect"] = nil
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusUnauthorized, redirect.(string))
	}
	sess.Values["redirect"] = nil
	sess.Values["userinfo"] = data
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusPermanentRedirect, redirect.(string))
}

func getAuths(c echo.Context) error {
	auths, err := services.GetAuths(c.Get("db").(*gorm.DB))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, auths)
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

func getUserDataFromGoogle(code string, oauthConfig *oauth2.Config, provider string) ([]byte, error) {
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req := auth.GetUserDataClient(provider, token)
	httpClinet := &http.Client{}
	response, err := httpClinet.Do(req)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return contents, nil
}

func addAuth(c echo.Context) error {
	auth := new(models.Auth)
	if err := c.Bind(auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(auth); err != nil {
		return err
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
	auth, err := services.GetAuthById(c.Get("db").(*gorm.DB), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, auth)
}

func replaceAuthById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	authFromDB, err := services.GetAuthById(db, id)
	if err != nil {
		return err
	}
	var auth models.Auth
	if err := c.Bind(&auth); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(auth); err != nil {
		return err
	}
	authFromDB.Authtype = auth.Authtype
	authFromDB.UserID = auth.UserID
	if err := services.ReplaceAuth(db, authFromDB); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
