package controllers

import (
	"apprit/store/api/v1/auth"
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/services"
	"apprit/store/api/v1/utils"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gorm.io/gorm"
)

func GetUsersGroup(e *echo.Group) {
	g := e.Group("/users")
	g.Use(middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.GET("", getUsers)
	g.POST("", addUser)
	g.DELETE("/:id", deleteUserById)
	g.GET("/:id", getUserById)
	g.PUT("/:id", replaceUserById)
}

func GetUserGroup(e *echo.Group) {
	g := e.Group("/user")
	g.Use(middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.DELETE("", deleteUserById)
	g.GET("/transactions", getUserTransactions)
	g.GET("/payments", getUserPayments)
	g.GET("/me", getUserData)
}

func getUserData(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return echo.ErrUnauthorized
	}
	claims := user.Claims.(*auth.JwtCustomClaims)
	userData, err := services.GetUserById(c.Get("db").(*gorm.DB), uint64(claims.ID))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, models.ConverUserToUserData(userData))
}

func getUserTransactions(c echo.Context) error {
	return c.JSON(http.StatusOK, []models.Transaction{})
}

func getUserPayments(c echo.Context) error {
	return c.JSON(http.StatusOK, []models.Payment{})
}

func getUsers(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return echo.ErrUnauthorized
	}
	claims := user.Claims.(*auth.JwtCustomClaims)
	if claims.Admin == false {
		return echo.ErrUnauthorized
	}
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
	userToken := c.Get("user").(*jwt.Token)
	if userToken == nil {
		user.Admin = false
	}
	claims := userToken.Claims.(*auth.JwtCustomClaims)
	if claims.Admin == false {
		user.Admin = false
	}
	_, err := services.AddUser(c.Get("db").(*gorm.DB), *user)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getUserById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	userToken := c.Get("user").(*jwt.Token)
	if userToken == nil {
		return echo.ErrUnauthorized
	}
	claims := userToken.Claims.(*auth.JwtCustomClaims)
	if claims.Admin == false || claims.ID != uint(id) {
		return echo.ErrUnauthorized
	}
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
	userToken := c.Get("user").(*jwt.Token)
	if userToken == nil {
		return echo.ErrUnauthorized
	}
	claims := userToken.Claims.(*auth.JwtCustomClaims)
	if claims.Admin == false || claims.ID != uint(id) {
		return echo.ErrUnauthorized
	}
	if err := services.DeleteUserById(c.Get("db").(*gorm.DB), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func replaceUserById(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	if userToken == nil {
		return echo.ErrUnauthorized
	}
	claims := userToken.Claims.(*auth.JwtCustomClaims)
	if claims.Admin == false {
		return echo.ErrUnauthorized
	}
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
