package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func initServe() *echo.Echo {
	e := echo.New()
	e.Static("tmp", "tmp")
	e.GET("/", handler)
	return e
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

func main() {
	serve := initServe()

	serve.Logger.Fatal(serve.Start(":8080"))
}