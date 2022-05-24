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

func GetPaymentGroup(e *echo.Group) {
	g := e.Group("/payments")
	g.GET("", getPayments, utils.CheckAutorization)
	g.POST("", addPayment, utils.CheckAutorization)
	g.DELETE("/:id", deletePaymentById, utils.CheckAutorization)
	g.GET("/:id", getPaymentById, utils.CheckAutorization)
	g.PUT("/:id", replacePaymentById, utils.CheckAutorization)
}

func getPayments(c echo.Context) error {
	payments, err := services.GetPayments(c.Get("db").(*gorm.DB))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, payments)
}

func addPayment(c echo.Context) error {
	payment := new(models.Payment)
	if err := utils.BindAndValidateObject(c, payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.AddPayment(c.Get("db").(*gorm.DB), *payment); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getPaymentById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	payment, err := services.GetPaymentById(c.Get("db").(*gorm.DB), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, payment)
}

func deletePaymentById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.DeletePaymentById(c.Get("db").(*gorm.DB), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func replacePaymentById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	var payment models.Payment
	if err := utils.BindAndValidateObject(c, &payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.ReplacePayment(db, id, payment); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
