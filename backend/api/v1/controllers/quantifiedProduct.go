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

func GetQuantifiedProductGroup(e *echo.Group) {
	g := e.Group("/quantifiedProducts")
	g.GET("", getQuantifiedProducts, utils.CheckAutorization)
	g.POST("", addQuantifiedProduct, utils.CheckAutorization)
	g.DELETE("/:id", deleteQuantifiedProductById, utils.CheckAutorization)
	g.GET("/:id", getQuantifiedProductById, utils.CheckAutorization)
	g.PUT("/:id", replaceQuantifiedProductById, utils.CheckAutorization)
}

func getQuantifiedProducts(c echo.Context) error {
	quantifiedProducts, err := services.GetQuantifiedProducts(c.Get("db").(*gorm.DB))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, quantifiedProducts)
}

func addQuantifiedProduct(c echo.Context) error {
	quantifiedProduct := new(models.QuantifiedProduct)
	if err := utils.BindAndValidateObject(c, quantifiedProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.AddQuantifiedProduct(c.Get("db").(*gorm.DB), *quantifiedProduct); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getQuantifiedProductById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	quantifiedProduct, err := services.GetQuantifiedProductById(c.Get("db").(*gorm.DB), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, quantifiedProduct)
}

func deleteQuantifiedProductById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.DeleteQuantifiedProductById(c.Get("db").(*gorm.DB), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func replaceQuantifiedProductById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	var quantifiedProduct models.QuantifiedProduct
	if err := utils.BindAndValidateObject(c, &quantifiedProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.ReplaceQuantifiedProduct(db, id, quantifiedProduct); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}