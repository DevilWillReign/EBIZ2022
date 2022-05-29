package controllers

import (
	"apprit/store/api/v1/auth"
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/services"
	"apprit/store/api/v1/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func GetCategoryGroup(e *echo.Group) {
	g := e.Group("/categories")
	g.GET("", getCategories)
	g.POST("", addCategory, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.DELETE("/:id", deleteCategoryById, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.GET("/:id", getCategoryById)
	g.PUT("/:id", replaceCategoryById, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
}

func getCategories(c echo.Context) error {
	categories, err := services.GetCategories(c.Get("db").(*gorm.DB))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, categories)
}

func addCategory(c echo.Context) error {
	category := new(models.Category)
	if err := utils.BindAndValidateObject(c, category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if _, err := services.AddCategory(c.Get("db").(*gorm.DB), *category); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getCategoryById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	category, err := services.GetCategoryById(c.Get("db").(*gorm.DB), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, category)
}

func deleteCategoryById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.DeleteCategoryById(c.Get("db").(*gorm.DB), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func replaceCategoryById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	var category models.Category
	if err := utils.BindAndValidateObject(c, &category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.ReplaceCategory(db, id, category); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
