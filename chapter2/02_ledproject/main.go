package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	// led := machine.LED
	led := machine.LCD_BACKLIGHT
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// TinyGoは基本的にmain関数を抜けないのでforで無限ループさせる
	for {
		fmt.Printf("Hello World\r")

		led.Low()
		time.Sleep(time.Millisecond * 1000)

		led.High()
		time.Sleep(time.Millisecond * 1000)

		// led.Toggle()
	}
}
