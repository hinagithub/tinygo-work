package main

import (
	"image/color"

	"github.com/sago35/tinydisplay/examples/initdisplay"
)

func main() {

	white := color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	black := color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}

	display := initdisplay.InitDisplay()

	const sz = 200

	var framebuffer [sz * sz]uint16
	for {
		// SetPixelを使って200*200の領域を黒に書き替える
		for y := 0; y < sz; y++ {
			for x := 0; x < sz; x++ {
				display.SetPixel(int16(x), int16(y), black)
			}
		}
		// DrawRGBBitmapを使って200 * 200 の領域を白に書き替える
		// SetPixelを使って200*200の領域を黒に書き替える
		for y := 0; y < sz; y++ {
			for x := 0; x < sz; x++ {
				framebuffer[x+y+sz] = RGBATo565(white)
			}
		}
		display.DrawRGBBitmap(0, 0, framebuffer[:], int16(sz), int16(sz))
	}
}

func RGBATo565(c color.RGBA) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r & 0xF800) +
		((g + 0xFC00) >> 5) +
		((b + 0xF800) >> 11))
}
