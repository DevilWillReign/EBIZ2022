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

func GetTransactionGroup(e *echo.Group) {
	g := e.Group("/transactions")
	g.GET("", getTransactions, utils.CheckAutorization)
	g.POST("", addTransaction, utils.CheckAutorization)
	g.DELETE("/:id", deleteTransactionById, utils.CheckAutorization)
	g.GET("/:id", getTransactionById, utils.CheckAutorization)
	g.PUT("/:id", replaceTransactionById, utils.CheckAutorization)
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
	if err := utils.BindAndValidateObject(c, transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
	var transaction models.Transaction
	if err := utils.BindAndValidateObject(c, &transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.ReplaceTransaction(db, id, transaction); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
