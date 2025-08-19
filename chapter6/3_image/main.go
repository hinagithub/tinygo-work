package main

import (
	"image/color"
	"image/jpeg"
	"log"
	"strings"

	"github.com/sago35/tinydisplay/examples/initdisplay"
)

func main() {
	display := initdisplay.InitDisplay()
	img, err := jpeg.Decode(strings.NewReader(tinygo_jpg))
	if err != nil {
		log.Fatal(err)
	}

	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < image.Bounds().Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			display.SetPixel(int16(x), int16(y), color.RGBA{
				R: uint8(r >> 8), G: unit8(g >> 8),
				B: uint8(b >> 8), A: unit8(0xFF),
			})
		}
	}
	select {}
}
