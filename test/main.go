package main

import (
	"fmt"
	"github.com/nickalie/go-webpbin"
	"os"
)
func main() {
	i, _ := os.Open("Go-Logo_LightBlue.png")
	o, _ := os.Create("New.webp")
	err := webpbin.NewCWebP().Quality(90).Input(i).Output(o).Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Completed")
}

/*
func main() {
	f, err := os.Open("Go-Logo_LightBlue.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	bou := img.Bounds()

	o, err := os.Create("New.webp")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(o)

	config, _ := webp.ConfigPreset(webp.PresetDefault, 90)
	m := image.NewNRGBA(image.Rect(0, 0, bou.Dx(), bou.Dy()))
	draw.Draw(m, m.Bounds(), img, bou.Min, draw.Src)
	if err = webp.EncodeRGBA(w, img, config); err != nil {
		panic(err)
	}

	w.Flush()
	o.Close()
}
 */