package utils

import (
	"apprit/store/api/v1/models"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/argon2"
	"golang.org/x/oauth2"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func GetEnv(name string, defaultval string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return defaultval
}

func BindAndValidateObject[E models.Entity](c echo.Context, entity *E) error {
	if err := c.Bind(&entity); err != nil {
		return err
	}
	return c.Validate(entity)
}

func GetSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	return sess
}

func CheckAutorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if _, err := c.Request().Cookie("admin"); err != nil {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetTokenCookie(token string) *http.Cookie {
	return &http.Cookie{Name: "userinfo", Value: token, HttpOnly: true, SameSite: http.SameSiteStrictMode, Path: "/"}
}

func HashPassword(password []byte, salt []byte) []byte {
	return argon2.IDKey(password, salt, 3, 32*1024, 4, 32)
}

type errorDetails struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func GetOauthErrorDetails(err error) {
	if rErr, ok := err.(*oauth2.RetrieveError); ok {
		details := new(errorDetails)
		if err := json.Unmarshal(rErr.Body, details); err != nil {
			log.Println(err)
		}

		log.Println(details.Error, details.ErrorDescription)
	}
}
