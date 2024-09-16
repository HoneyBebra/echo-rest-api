package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"service/src/api/v1"
	"service/src/core"
	"strings"
	"time"
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

	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutMinute) * time.Minute,
		WriteTimeout: time.Duration(cfg.WriteTimeoutMinute) * time.Minute,
	}
	return e.StartServer(s)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
