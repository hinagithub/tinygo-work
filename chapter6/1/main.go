package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ili9341"
	// ""
)

func main() {
	// バックライトをON（これが重要でした！）
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
	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	pink := color.RGBA{R: 255, G: 192, B: 203, A: 255}

	// 黒い背景
	display.FillScreen(black)

	// 白い三角形を描画
	for y := int16(10); y < 100; y++ {
		for x := int16(10); x <= 100; x++ {
			if x < y {
				display.SetPixel(x, y, white)
			}
		}
	}

	// display.SetPixel(230, 30, pink)
	// display.SetPixel(230, 31, pink)
	// display.SetPixel(230, 32, pink)
	// display.SetPixel(230, 33, pink)

	// ハートの中心座標
	centerX := int16(160)
	centerY := int16(120)

	// ハートの形を描画
	for y := int16(-10); y <= 6; y++ {
		for x := int16(-15); x <= 15; x++ {
			// ハート形の方程式を使用
			// 左の円
			leftCircle := (x+5)*(x+5) + (y-2)*(y-2)
			// 右の円
			rightCircle := (x-5)*(x-5) + (y-2)*(y-2)
			// 下の三角部分
			triangle := (x*x+2*y+2*y*y >= 0)

			if (leftCircle <= 25 || rightCircle <= 25) && triangle {
				display.SetPixel(centerX+x, centerY+y, pink)
			}
		}
	}

	select {}
}

// ピンクのハートを描画
func drawHeart() {
	// ハートの中心座標
	centerX := int16(160)
	centerY := int16(120)

	// ハートの形を描画
	for y := int16(-10); y <= 6; y++ {
		for x := int16(-15); x <= 15; x++ {
			// ハート形の方程式を使用
			// 左の円
			leftCircle := (x+5)*(x+5) + (y-2)*(y-2)
			// 右の円
			rightCircle := (x-5)*(x-5) + (y-2)*(y-2)
			// 下の三角部分
			triangle := (x*x+2*y+2*y*y >= 0)

			if (leftCircle <= 25 || rightCircle <= 25) && triangle {
				display.SetPixel(centerX+x, centerY+y, pink)
			}
		}
	}
}
