package main

import (
	"image/color"
	"machine"
	"math"
	"time"

	"tinygo.org/x/drivers/ili9341"
)

func main() {
	backlight := machine.LCD_BACKLIGHT
	backlight.Configure(machine.PinConfig{Mode: machine.PinOutput})
	backlight.High()

	time.Sleep(100 * time.Millisecond)

	// SPI設定
	machine.SPI3.Configure(machine.SPIConfig{
		SCK:       machine.LCD_SCK_PIN,
		SDO:       machine.LCD_SDO_PIN,
		SDI:       machine.LCD_SDI_PIN,
		Frequency: 40000000,
	})

	display := ili9341.NewSPI(
		machine.SPI3,
		machine.LCD_DC,
		machine.LCD_SS_PIN,
		machine.LCD_RESET,
	)

	display.Configure(ili9341.Config{})
	time.Sleep(100 * time.Millisecond)

	// 色定義
	// black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	pink := color.RGBA{R: 255, G: 192, B: 203, A: 255}
	blue := color.RGBA{R: 173, G: 216, B: 230, A: 255}
	yellow := color.RGBA{R: 255, G: 255, B: 150, A: 255}
	purple := color.RGBA{R: 221, G: 160, B: 221, A: 255}
	// orange := color.RGBA{R: 255, G: 204, B: 153, A: 255}
	// green := color.RGBA{R: 144, G: 238, B: 144, A: 255}

	display.FillScreen(white)

	// 四角を書く
	for i := int16(0); i < 100; i++ {
		for j := int16(0); j < 200; j++ {
			display.SetPixel(i, j, yellow)
		}
	}

	// 縦線を引く
	for x := int16(0); x < 240; x++ {
		display.SetPixel(x, 160, purple)
	}

	// 横線を引く
	for y := int16(0); y < 320; y++ {
		display.SetPixel(120, y, blue)
	}

	// ハートを書く
	centerX := float32(160)
	centerY := float32(120)
	size := float32(5)
	for t := float32(0); t <= 2*3.14159; t += 0.01 {
		x := size * 16 * float32(math.Pow(math.Sin(float64(t)), 3))
		y := -(size*13*float32(math.Cos(float64(t))) - size*5*float32(math.Cos(2*float64(t))) - size*2*float32(math.Cos(3*float64(t))) - float32(math.Cos(4*float64(t))))
		display.SetPixel(int16(centerY+y), int16(centerX+x), pink)
	}
	select {}
}
