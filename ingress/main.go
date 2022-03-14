package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	hostname, _ := os.Hostname()

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, This is : " + hostname + "\n")
	})

	e.GET("/welcome/test", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Welcome!, This is : " + hostname + "\n")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
