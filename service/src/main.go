package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"service/src/api/v1"
	"service/src/core"
	"strings"
)

func run() error {
	cfg := core.GetConfig()

	e := echo.New()

	if strings.ToLower(cfg.Debug) == "true" {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	api := e.Group(cfg.ApiPrefix)

	// Test handler
	api.GET("/", v1.GetHello)

	return e.Start(fmt.Sprintf("%s:%s", "0.0.0.0", cfg.Port))
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
