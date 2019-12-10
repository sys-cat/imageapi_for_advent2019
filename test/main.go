package main

import (
	"fmt"
	"image"
	"golang.org/x/image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	// JPEG
	j, err := os.Open("../tmp/Go-Logo_LightBlue.jpg")
	if err != nil {
		panic(err)
	}
	defer j.Close()
	img, err := jpeg.Decode(j)
	if err != nil {
		panic(err)
	}
	bound := img.Bounds()
	fmt.Printf("Before\tGo-Logo_LightBlue.jpg\tX : %d px, Y : %d px\n", bound.Dx(), bound.Dy())

	dst := image.NewRGBA(image.Rect(0, 0, bound.Dx()/2, bound.Dy()/2))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, bound, draw.Over, nil)

	jOutput, err := os.Create("./new.jpg")
	if err != nil {
		panic(err)
	}
	defer jOutput.Close()

	err = jpeg.Encode(jOutput, dst, &jpeg.Options{Quality: 100})
	if err != nil {
		panic(err)
	}

	reopen, err := os.Open("./new.jpg")
	if err != nil {
		panic(err)
	}
	defer reopen.Close()
	img, _ = jpeg.Decode(reopen)
	bound = img.Bounds()
	fmt.Printf("After\tGo-Logo_LightBlue.jpg\tX : %d px, Y : %d px\n", bound.Dx(), bound.Dy())

	// PNG
	p, err := os.Open("../tmp/Go-Logo_Yellow.png")
	if err != nil {
		panic(err)
	}
	defer p.Close()
	img, err = png.Decode(p)
	if err != nil {
		panic(err)
	}
	bound = img.Bounds()
	fmt.Printf("Before\tGo-Logo_Yellow.png\tX : %d px, Y : %d px\n", bound.Dx(), bound.Dy())

	dst = image.NewRGBA(image.Rect(0, 0, bound.Dx()/2, bound.Dy()/2))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, bound, draw.Over, nil)

	pOutput, err := os.Create("./new.png")
	if err != nil {
		panic(err)
	}
	defer pOutput.Close()

	err = png.Encode(pOutput, dst)
	if err != nil {
		panic(err)
	}

	reopen, err = os.Open("./new.png")
	if err != nil {
		panic(err)
	}
	defer reopen.Close()
	img, _ = png.Decode(reopen)
	bound = img.Bounds()
	fmt.Printf("After\tGo-Logo_Yellow.png\tX : %d px, Y : %d px\n", bound.Dx(), bound.Dy())
}