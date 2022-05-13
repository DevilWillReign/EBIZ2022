package controllers

import (
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetCategoryGroup(e *echo.Group) {
	g := e.Group("/categories")
	g.GET("", getCategories)
	g.POST("", addCategory)
	g.DELETE("/:id", deleteCategoryById)
	g.GET("/:id", getCategoryById)
	g.PUT("/:id", replaceCategoryById)
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
	if err := c.Bind(category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(category); err != nil {
		return err
	}
	if err := services.AddCategory(c.Get("db").(*gorm.DB), *category); err != nil {
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
	categoryFromDB, err := services.GetCategoryById(db, id)
	if err != nil {
		return err
	}
	var category models.Category
	if err := c.Bind(&category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(category); err != nil {
		return err
	}
	categoryFromDB.Name = category.Name
	if err := services.ReplaceCategory(db, categoryFromDB); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
