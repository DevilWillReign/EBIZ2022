package controllers

import (
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetPaymentGroup(e *echo.Group) {
	g := e.Group("/payments")
	g.GET("", getPayments)
	g.POST("", addPayment)
	g.DELETE("/:id", deletePaymentById)
	g.GET("/:id", getPaymentById)
	g.PUT("/:id", replacePaymentById)
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
	if err := c.Bind(payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(payment); err != nil {
		return err
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
	paymentFromDB, err := services.GetPaymentById(db, id)
	if err != nil {
		return err
	}
	var payment models.Payment
	if err := c.Bind(&payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(payment); err != nil {
		return err
	}
	paymentFromDB.Total = payment.Total
	paymentFromDB.TransactionID = payment.TransactionID
	if err := services.ReplacePayment(db, paymentFromDB); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
