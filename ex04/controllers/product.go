package controllers

import (
	"apprit/store/models"
	"apprit/store/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetProductGroup(e *echo.Echo) {
	g := e.Group("/products")
	g.GET("", getProducts)
	g.POST("", addProduct)
	g.DELETE("/:id", deleteProductById)
	g.GET("/:id", getProductById)
	g.PUT("/:id", replaceProductById)
}

func getProducts(c echo.Context) error {
	products, err := services.GetProducts(c.Get("db").(*gorm.DB))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}

func addProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(product); err != nil {
		return err
	}
	if err := services.AddProduct(c.Get("db").(*gorm.DB), *product); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getProductById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	product, err := services.GetProductById(c.Get("db").(*gorm.DB), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, product)
}

func deleteProductById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.DeleteProductById(c.Get("db").(*gorm.DB), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func replaceProductById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	productFromDB, err := services.GetProductById(db, id)
	if err != nil {
		return err
	}
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(product); err != nil {
		return err
	}
	productFromDB.Name = product.Name
	productFromDB.Code = product.Code
	productFromDB.Price = product.Price
	productFromDB.CategoryID = product.CategoryID
	if err := services.ReplaceProduct(db, productFromDB); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
