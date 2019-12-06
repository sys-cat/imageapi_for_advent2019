package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/echo"
	"image"
	"net/http"
	"os"
	"image/jpeg"
	"github.com/harukasan/go-libwebp/webp"
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
	w := bufio.NewWriter(o)


	config, _ := webp.ConfigPreset(webp.PresetDefault, 90)
	if err = webp.EncodeRGBA(w, image.NewRGBA(image.Rect(0, 0, bou.Dx(), bou.Dy())), config);err != nil {
		return c.JSON(http.StatusInternalServerError, &Error{Message:"convert image is failed"})
	}

	img, _, err = image.Decode(o)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Error{Message:"cant get converted image"})
	}
	bouN := img.Bounds()

	w.Flush()
	o.Close()
	f.Close()

	return c.String(http.StatusOK, fmt.Sprintf("jpeg image X size : %d, Y size: %d\nwebp image X size : %d, Y size : %c\n", bou.Dx(), bou.Dy(), bouN.Dx(), bouN.Dy()))
}

func main() {
	serve := initServe()

	serve.Logger.Fatal(serve.Start(":8080"))
}