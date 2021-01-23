package main

import (
	"net/http"
	"io"
	"html/template"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		return c.String(http.StatusInternalServerError, "Fatal error invoked.")
	})


	e.Logger.Fatal(e.Start(":8080"))
}