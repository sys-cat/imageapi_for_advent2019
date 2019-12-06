package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/echo"
	"image"
	"net/http"
	"os"
	_ "image/jpeg"
	"github.com/sys-cat/go-libwebp/webp"
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
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return c.JSON(http.StatusGone, &Error{Message:"File is gone"})
	}
	bou := img.Bounds()

	newimg := image.NewNRGBA(image.Rect(0, 0, bou.Dx(), bou.Dy()))
	o, err := os.Create("tmp/New.webp")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Error{Message:"create new image is failed"})
	}
	w := bufio.NewWriter(o)
	defer func() {
		w.Flush()
		o.Close()
	}()

	config, _ := webp.ConfigPreset(webp.PresetDefault, 90)
	if err = webp.EncodeRGBA(w, newimg, config);err != nil {
		return c.JSON(http.StatusInternalServerError, &Error{Message:"convert image is failed"})
	}

	img, _, err = image.Decode(o)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Error{Message:"cant get converted image"})
	}
	bouN := img.Bounds()

	return c.String(http.StatusOK, fmt.Sprintf("jpeg image X size : %d, Y size: %d\nwebp image X size : %d, Y size : %c\n", bou.Dx(), bou.Dy(), bouN.Dx(), bouN.Dy()))
}

func main() {
	serve := initServe()

	serve.Logger.Fatal(serve.Start(":8080"))
}