package utils

import (
	"apprit/store/api/v1/models"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
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
