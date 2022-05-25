package controllers

import (
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/services"
	"apprit/store/api/v1/utils"
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
	g.GET("/:id/transactions", getUserTransactions)
}

func getUserTransactions(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Transaction{})
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
	if err := utils.BindAndValidateObject(c, user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.AddUser(c.Get("db").(*gorm.DB), *user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, models.UserData{Name: user.Username, Email: user.Email})
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
	var user models.User
	if err := utils.BindAndValidateObject(c, &user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.ReplaceUser(db, id, user); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
