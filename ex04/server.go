package main

import (
	"apprit/store/controllers"
	"apprit/store/database"
	"apprit/store/utils"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := database.CreateDatabase()
	database.InitializeDatabaseData(db)
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	controllers.GetCategoryGroup(e)
	controllers.GetAuthGroup(e)
	controllers.GetPaymentGroup(e)
	controllers.GetProductGroup(e)
	controllers.GetQuantifiedProductGroup(e)
	controllers.GetTransactionGroup(e)
	controllers.GetUserGroup(e)
	e.Logger.Fatal(e.Start(":" + utils.GetEnv("PORT", "9000")))
}
