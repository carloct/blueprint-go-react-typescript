package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	RegisterFrontendHandlers(e)

	api := e.Group("/api")
	{
		api.GET("/hello", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello from API!")
		})
	}

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	e.Logger.Fatal(e.Start(":8000"))
}
