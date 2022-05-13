package controllers

import (
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUserGroup(e *echo.Group) {
	g := e.Group("/users")
	g.GET("", getUsers)
	g.POST("", addUser)
	g.DELETE("/:id", deleteUserById)
	g.GET("/:id", getUserById)
	g.PUT("/:id", replaceUserById)
}

func getUsers(c echo.Context) error {
	users, err := services.GetUsers(c.Get("db").(*gorm.DB))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func addUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return err
	}
	if err := services.AddUser(c.Get("db").(*gorm.DB), *user); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getUserById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := services.GetUserById(c.Get("db").(*gorm.DB), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func deleteUserById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.DeleteUserById(c.Get("db").(*gorm.DB), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func replaceUserById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	userFromDB, err := services.GetUserById(db, id)
	if err != nil {
		return err
	}
	var user models.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return err
	}
	userFromDB.Username = user.Username
	userFromDB.Email = user.Email
	userFromDB.Password = user.Password
	if err := services.ReplaceUser(db, userFromDB); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
