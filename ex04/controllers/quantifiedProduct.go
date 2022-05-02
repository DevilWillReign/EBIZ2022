package controllers

import (
	"apprit/store/models"
	"apprit/store/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetQuantifiedProductGroup(e *echo.Echo) {
	g := e.Group("/quantifiedProducts")
	g.GET("", getQuantifiedProducts)
	g.POST("", addQuantifiedProduct)
	g.DELETE("/:id", deleteQuantifiedProductById)
	g.GET("/:id", getQuantifiedProductById)
	g.PUT("/:id", replaceQuantifiedProductById)
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
	if err := c.Bind(quantifiedProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(quantifiedProduct); err != nil {
		return err
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
	quantifiedProductFromDB, err := services.GetQuantifiedProductById(db, id)
	if err != nil {
		return err
	}
	var quantifiedProduct models.QuantifiedProduct
	if err := c.Bind(&quantifiedProduct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(quantifiedProduct); err != nil {
		return err
	}
	quantifiedProductFromDB.Name = quantifiedProduct.Name
	quantifiedProductFromDB.Price = quantifiedProduct.Price
	quantifiedProductFromDB.Code = quantifiedProduct.Code
	quantifiedProductFromDB.Quantity = quantifiedProduct.Quantity
	quantifiedProductFromDB.CategoryID = quantifiedProduct.CategoryID
	quantifiedProductFromDB.TransactionID = quantifiedProduct.TransactionID
	if err := services.ReplaceQuantifiedProduct(db, quantifiedProductFromDB); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
