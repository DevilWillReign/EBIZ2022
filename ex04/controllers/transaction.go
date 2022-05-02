package controllers

import (
	"apprit/store/models"
	"apprit/store/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetTransactionGroup(e *echo.Echo) {
	g := e.Group("/transactions")
	g.GET("", getTransactions)
	g.POST("", addTransaction)
	g.DELETE("/:id", deleteTransactionById)
	g.GET("/:id", getTransactionById)
	g.PUT("/:id", replaceTransactionById)
}

func getTransactions(c echo.Context) error {
	transactions, err := services.GetTransactions(c.Get("db").(*gorm.DB))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, transactions)
}

func addTransaction(c echo.Context) error {
	transaction := new(models.Transaction)
	if err := c.Bind(transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(transaction); err != nil {
		return err
	}
	if err := services.AddTransaction(c.Get("db").(*gorm.DB), *transaction); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getTransactionById(c echo.Context) error {
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	transactionFromDB, err := services.GetTransactionById(db, id)
	if err != nil {
		return err
	}
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(transaction); err != nil {
		return err
	}
	transactionFromDB.UserID = transaction.UserID
	if err := services.ReplaceTransaction(db, transactionFromDB); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
