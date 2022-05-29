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

func GetProductGroup(e *echo.Group) {
	g := e.Group("/products")
	g.GET("", getProducts)
	g.POST("", addProduct, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.DELETE("/:id", deleteProductById, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.GET("/:id", getProductById)
	g.PUT("/:id", replaceProductById, middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
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
	if err := utils.BindAndValidateObject(c, product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if _, err := services.AddProduct(c.Get("db").(*gorm.DB), *product); err != nil {
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
	var product models.Product
	if err := utils.BindAndValidateObject(c, &product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.ReplaceProduct(db, id, product); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
