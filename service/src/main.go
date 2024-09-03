package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"service/src/api/v1"
)

func run() error {
	e := echo.New()

	e.Use(middleware.Logger())

	api := e.Group("/api/v1")

	// Test handler
	api.GET("/", v1.GetHello)

	return e.Start("0.0.0.0:8000")
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
