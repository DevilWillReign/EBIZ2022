package controllers

import (
	"apprit/store/models"
	"apprit/store/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAuthGroup(e *echo.Echo) {
	g := e.Group("/auths")
	g.GET("", getAuths)
	g.POST("", addAuth)
	g.DELETE("/:id", deleteAuthById)
	g.GET("/:id", getAuthById)
	g.PUT("/:id", replaceAuthById)
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
