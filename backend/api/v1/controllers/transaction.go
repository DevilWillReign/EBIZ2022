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

func GetTransactionGroup(e *echo.Group) {
	g := e.Group("/transactions")
	g.Use(middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.GET("", getTransactions)
	g.POST("", addTransaction)
	g.DELETE("/:id", deleteTransactionById)
	g.GET("/:id", getTransactionById)
	g.PUT("/:id", replaceTransactionById)
	g.GET("/:id/extended", getExtendedTransactionById)
}

func getExtendedTransactionById(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	transaction, err := services.GetTransactionById(db, id)
	if err != nil {
		return err
	}
	payment, err := services.GetPaymentByTransactionId(db, uint64(transaction.ID))
	if err != nil {
		return err
	}
	quantifiedProducts, err := services.GetQuantifiedProductByTransactionId(db, uint64(transaction.ID))
	if err != nil {
		return err
	}
	transaction.Payment = payment
	transaction.QuantifiedProducts = quantifiedProducts
	return c.JSON(http.StatusOK, transaction)
}

func getTransactions(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	transactions, err := services.GetTransactions(c.Get("db").(*gorm.DB))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, models.ResponseArrayEntity[models.Transaction]{Elements: transactions})
}

func addTransaction(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	transaction := models.PostTransaction{}
	if err := utils.BindAndValidateObject(c, &transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if _, err := services.AddTransaction(c.Get("db").(*gorm.DB), transaction); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getTransactionById(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	transaction, err := services.GetTransactionById(c.Get("db").(*gorm.DB), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, transaction)
}

func deleteTransactionById(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.DeleteTransactionById(c.Get("db").(*gorm.DB), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func replaceTransactionById(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	transaction := models.Transaction{}
	if err := utils.BindAndValidateObject(c, &transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.ReplaceTransaction(db, id, transaction); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
