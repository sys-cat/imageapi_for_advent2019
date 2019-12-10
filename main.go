package main

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
	"image/jpeg"
)

type (
	Error struct {
		Message string
	}
)

func initServe() *echo.Echo {
	e := echo.New()
	e.Static("tmp", "tmp")
	e.GET("/", handler)
	return e
}

func handler(c echo.Context) error {
	f, err := os.Open("tmp/Go-Logo_LightBlue.jpg")
	if err != nil {
		return c.JSON(http.StatusNotFound, &Error{
			Message:"file not found",
		})
	}

	img, err := jpeg.Decode(f)
	if err != nil {
		return c.JSON(http.StatusGone, &Error{Message:"File is gone"})
	}
	bou := img.Bounds()

	o, err := os.Create("tmp/New.webp")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Error{Message:"create new image is failed"})
	}

	defer func() {
		o.Close()
		f.Close()
	}()

}

func main() {
	serve := initServe()

	serve.Logger.Fatal(serve.Start(":8080"))
}