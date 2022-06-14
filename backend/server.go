package main

import (
	"apprit/store/api/v1/controllers"
	"apprit/store/api/v1/database"
	"apprit/store/api/v1/utils"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	/**
	DB initialization
	*/
	db := database.CreateDatabase()
	database.InitializeDatabaseData(db)
	/**
	Echo initialization
	*/
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(utils.GetEnv("FRONT_HOST", "http://localhost:9001"), ","),
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
	}))
	// Session
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(utils.GetEnv("SESSION_SECRET", "secret")))))
	// Adding validator
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	// Adding db to echo context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})
	// Adding api controllers
	api := e.Group("/api/v1")
	controllers.GetCategoryGroup(api)
	controllers.GetAuthGroup(api)
	controllers.GetPaymentGroup(api)
	controllers.GetProductGroup(api)
	controllers.GetQuantifiedProductGroup(api)
	controllers.GetTransactionGroup(api)
	controllers.GetUsersGroup(api)
	controllers.GetUserGroup(api)
	e.Logger.Fatal(e.Start(utils.GetEnv("API_HOST", "") + ":" + utils.GetEnv("API_PORT", "9000")))
}
