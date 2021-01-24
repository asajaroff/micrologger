package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"net/http"
	"os"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Micrologger v0.0.1")
	})

	e.GET("/fatal", func(c echo.Context) error {
		e.Logger.Fatal("This is the sample fatal output. Designates very severe error events that will presumably lead the application to abort." )
		e.Logger.Debug("The /fatal endpoint has been requested. This causes the process to stop on purpose.")
		os.Exit(1)
		return c.String(http.StatusInternalServerError, "Endpoint requested: /fatal.")
	})

	e.GET("/error", func(c echo.Context) error {
		e.Logger.Error("This is a the sample error output. Designates error events that might still allow the application to continue running." )
		return c.String(http.StatusConflict, "Endpoint requested: /error")
	})

	e.GET("/warning", func(c echo.Context) error {
		e.Logger.Warn("This is the sample warning output. Designates potentially harmful situations")
		return c.String(http.StatusBadRequest, "Endpoint requested: /warning")
	})

	e.GET("/info", func(c echo.Context) error {
		e.Logger.Info("This is the sample info output. Designates informational messages that highlight the progress of the application at coarse-grained level.")
		return c.String(http.StatusOK, "Endpoint requested: /info")
	})

	e.GET("/debug", func(c echo.Context) error {
		e.Logger.Debug("This is the sample debug output. Designates fine-grained informational events that are most useful to debug an application.")
		return c.String(http.StatusOK, "Endpoint requested: /info")
	})

	e.Logger.Fatal(e.Start(":8080"))
}