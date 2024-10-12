package main

import (
	"embed"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	//go:embed frontend/dist/*
	dist embed.FS

	//go:embed frontend/dist/index.html
	indexHTML embed.FS

	distDirFS     = echo.MustSubFS(dist, "frontend/dist")
	distIndexHTML = echo.MustSubFS(indexHTML, "frontend/dist")
)

func RegisterFrontendHandlers(e *echo.Echo) {
	if os.Getenv("APP_ENV") == "development" {
		log.Println("Serving running in development mode")
		setupDevProxy(e)
		return
	}

	fmt.Printf("distIndexHTML: %v\n", distIndexHTML)

	e.FileFS("/", "index.html", distIndexHTML)
	e.StaticFS("/", distDirFS)
}

func setupDevProxy(e *echo.Echo) {
	// setup dev proxy
	url, err := url.Parse("http://localhost:5173")
	if err != nil {
		log.Fatal(err)
	}

	balancer := middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: url,
		},
	})

	e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: balancer,
		Skipper: func(c echo.Context) bool {
			return len(c.Path()) >= 4 && c.Path()[:4] == "/api"
		},
	}))
}
