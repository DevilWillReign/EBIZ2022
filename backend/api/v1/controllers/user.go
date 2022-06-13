package controllers

import (
	"apprit/store/api/v1/auth"
	"apprit/store/api/v1/models"
	"apprit/store/api/v1/services"
	"apprit/store/api/v1/utils"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shopspring/decimal"

	"gorm.io/gorm"
)

func GetUsersGroup(e *echo.Group) {
	g := e.Group("/users")
	g.Use(middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.GET("", getUsers)
	g.POST("", addUser)
	g.DELETE("/:id", deleteUserById)
	g.GET("/:id", getUserById)
	g.PUT("/:id", replaceUserById)
}

func GetUserGroup(e *echo.Group) {
	g := e.Group("/user")
	g.Use(middleware.JWTWithConfig(auth.GetCustomClaimsConfig()))
	g.GET("/transactions", getUserTransactions)
	g.POST("/transactions", addUserTransaction)
	g.GET("/payments", getUserPayments)
	g.POST("/payments", addUserPayments)
	g.GET("/transactions/:id", getUserTransactionById)
	g.GET("/transactions/:id/total", getUserTransactionTotalById)
	g.GET("/transactions/:transactionid/payment/:id", getUserPaymentsById)
	g.GET("/me", getUserData)
}

func getUserData(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return echo.ErrUnauthorized
	}
	claims := user.Claims.(*auth.JwtCustomClaims)
	userData, err := services.GetUserById(c.Get("db").(*gorm.DB), uint64(claims.ID))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, models.ConverUserToUserData(userData))
}

func getUserTransactions(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return echo.ErrUnauthorized
	}
	claims := user.Claims.(*auth.JwtCustomClaims)
	transactions, err := services.GetUserTransactions(c.Get("db").(*gorm.DB), uint64(claims.ID))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, models.ResponseArrayEntity[models.Transaction]{Elements: transactions})
}

func getUserTransactionById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return echo.ErrUnauthorized
	}
	claims := user.Claims.(*auth.JwtCustomClaims)
	db := c.Get("db").(*gorm.DB)
	transaction, err := services.GetUserTransactionById(db, uint64(claims.ID), id)
	if err != nil {
		return echo.ErrNotFound
	}
	if payment, err := services.GetPaymentByTransactionId(db, uint64(transaction.ID)); err != nil {
		transaction.Payment = models.Payment{}
	} else {
		transaction.Payment = payment
	}
	if quantifiedProducts, err := services.GetQuantifiedProductByTransactionId(db, uint64(transaction.ID)); err != nil {
		transaction.QuantifiedProducts = []models.QuantifiedProduct{}
	} else {
		transaction.QuantifiedProducts = quantifiedProducts
	}
	return c.JSON(http.StatusOK, transaction)
}

func getUserTransactionTotalById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return echo.ErrUnauthorized
	}
	claims := user.Claims.(*auth.JwtCustomClaims)
	db := c.Get("db").(*gorm.DB)
	transaction, err := services.GetUserTransactionById(db, uint64(claims.ID), id)
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, map[string]decimal.Decimal{"total": transaction.Total})
}

func getUserPaymentsById(c echo.Context) error {
	transactionid, err := strconv.ParseUint(c.Param("transactionid"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return echo.ErrUnauthorized
	}
	claims := user.Claims.(*auth.JwtCustomClaims)
	payment, err := services.GetUserPaymentById(c.Get("db").(*gorm.DB), uint64(claims.ID), transactionid, id)
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, payment)
}

func addUserTransaction(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return echo.ErrUnauthorized
	}
	claims := user.Claims.(*auth.JwtCustomClaims)
	transaction := models.UserTransaction{}
	if err := c.Bind(&transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	transaction.UserID = claims.ID
	total := decimal.NewFromInt(0)
	for _, product := range transaction.QuantifiedProducts {
		total = total.Add(product.Price.Mul(decimal.NewFromInt(int64(product.Quantity))))
	}
	transaction.Total = total
	if err := c.Validate(&transaction); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if auth.IsNotAdminAndSameUser(c, uint64(transaction.UserID)) {
		return echo.ErrUnauthorized
	}
	transactionid, err := services.AddUserTransaction(c.Get("db").(*gorm.DB), transaction)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, map[string]int64{"transactionid": transactionid})
}

func addUserPayments(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return echo.ErrUnauthorized
	}
	claims := user.Claims.(*auth.JwtCustomClaims)
	payment := models.Payment{}
	if err := utils.BindAndValidateObject(c, &payment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	if _, err := services.GetUserTransactionById(db, uint64(claims.ID), uint64(payment.TransactionID)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if _, err := services.AddPayment(db, payment); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func getUserPayments(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	if user == nil {
		return echo.ErrUnauthorized
	}
	claims := user.Claims.(*auth.JwtCustomClaims)
	transactions, err := services.GetUserTransactions(c.Get("db").(*gorm.DB), uint64(claims.ID))
	if err != nil {
		return echo.ErrNotFound
	}
	payments := []models.Payment{}
	for _, transaction := range transactions {
		payment, err := services.GetPaymentByTransactionId(c.Get("db").(*gorm.DB), uint64(transaction.ID))
		if err == nil {
			payments = append(payments, payment)
		}
	}
	return c.JSON(http.StatusOK, models.ResponseArrayEntity[models.Payment]{Elements: payments})
}

func getUsers(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	users, err := services.GetUsers(c.Get("db").(*gorm.DB))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, models.ResponseArrayEntity[models.User]{Elements: users})
}

func addUser(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	user := models.User{}
	if err := utils.BindAndValidateObject(c, &user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, err := services.AddUser(c.Get("db").(*gorm.DB), user)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func getUserById(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := services.GetUserById(c.Get("db").(*gorm.DB), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func deleteUserById(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.DeleteUserById(c.Get("db").(*gorm.DB), id); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func replaceUserById(c echo.Context) error {
	if auth.IsNotAdmin(c) {
		return echo.ErrUnauthorized
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db := c.Get("db").(*gorm.DB)
	user := models.User{}
	if err := utils.BindAndValidateObject(c, &user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.ReplaceUser(db, id, user); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
