package main

import (
	"fmt"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

var (
	ssid     = "tullys_Wi-Fi"
	password = ""
	serverIP = "127.0.0.1"
	port     = 8080
)

func main() {
	// LCD初期化
	machine.SPI3.Configure(machine.SPIConfig{
		SCK:       machine.LCD_SCK_PIN,
		SDO:       machine.LCD_SDO_PIN,
		SDI:       machine.LCD_SDI_PIN,
		Frequency: 48000000,
	})
	backlight := machine.LCD_BACKLIGHT
	backlight.Configure(machine.PinConfig{Mode: machine.PinOutput})
	display := ili9341.NewSPI(
		machine.SPI3,
		machine.LCD_DC,
		machine.LCD_SS_PIN,
		machine.LCD_RESET,
	)
	display.Configure(ili9341.Config{Rotation: ili9341.Rotation270})
	display.FillScreen(color.RGBA{0, 0, 0, 255})
	backlight.High()

	white := color.RGBA{255, 255, 255, 255}
	tinyfont.WriteLine(display, &freemono.Regular12pt7b, 10, 30, "Wi-Fi SSID: "+ssid, white)
	tinyfont.WriteLine(display, &freemono.Regular12pt7b, 10, 60, "TCP Server: "+serverIP, white)
	tinyfont.WriteLine(display, &freemono.Regular12pt7b, 10, 90, "Port: "+fmt.Sprintf("%d", port), white)
	tinyfont.WriteLine(display, &freemono.Regular12pt7b, 10, 120, "Send: hello from tinygo", white)
	tinyfont.WriteLine(display, &freemono.Regular12pt7b, 10, 150, "※通信部分は未実装", white)

	// ここにRTL8720DN制御によるTCP通信実装を追加予定

	for {
		time.Sleep(time.Second)
	}
}
