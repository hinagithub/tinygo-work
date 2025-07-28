package main

import (
	"image/color"

	"github.com/sago35/tinydisplay/examples/initdisplay"
)

func main() {
	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	display := initdisplay.InitDisplay()
	display.FillScreen(black)

	for y := int16(10); y < 100; y++ {
		for x := int16(10); x <= 100; x++ {
			if x < y {
				display.SetPixel(x, y, white)
			}
		}
	}
	select {}
}
