// P218~ これまでのまとめ
package main

import (
	"image/color"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"

	"github.com/sago35/tinydisplay/examples/initdisplay"
)

func main() {
	// black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	purple := color.RGBA{R: 218, G: 193, B: 255, A: 255}

	display := initdisplay.InitDisplay()
	display.FillScreen(purple)
	display.FillScreen(purple)
	tinydraw.Line(display, 10, 10, 100, 100, white)
	tinydraw.Rectangle(display, 110, 10, 90, 50, white)
	tinydraw.Circle(display, 40, 110, 30, white)
	tinydraw.Triangle(display, 110, 110, 200, 150, 150, 200, white)
	tinyfont.WriteLine(display, &freesans.Regular18pt7b, 100, 100, "Hello TinyGo", white)
	select {}
}
