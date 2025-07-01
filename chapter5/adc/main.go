// ADC = Analog to Digital Converter

package main

import (
	"fmt"
	"machine"
)

func main() {
	machine.InitADC()
	sensor := machine.ADC{Pin: machine.WIO_LIGHT}
	sensor.Configure(machine.ADCConfig{})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		val := sensor.Get()
		// 暗いとき　AD値が0x8000より小さい(1.65V以下)ときにLEDを点灯する
		if val < 0x8000 {
			led.High()
		} else {
			led.Low()
		}

		switch {
		case val < 0x1999: // 約 0 - 6553 (1/10)
			// 段階 1: 最も暗い
			led.High()
			fmt.Println("Brightness: Level 1 (Lowest)")
		case val < 0x3333: // 約 6554 - 13107 (2/10)
			// 段階 2
			fmt.Println("Brightness: Level 2")
		case val < 0x4CCC: // 約 13108 - 19660 (3/10)
			// 段階 3
			fmt.Println("Brightness: Level 3")
		case val < 0x6666: // 約 19661 - 26214 (4/10)
			// 段階 4
			fmt.Println("Brightness: Level 4")
		case val < 0x8000: // 約 26215 - 32768 (5/10)
			// 段階 5
			fmt.Println("Brightness: Level 5")
		case val < 0x9999: // 約 32769 - 39321 (6/10)
			// 段階 6
			fmt.Println("Brightness: Level 6")
		case val < 0xB333: // 約 39322 - 45874 (7/10)
			// 段階 7
			fmt.Println("Brightness: Level 7")
		case val < 0xCCCC: // 約 45875 - 52428 (8/10)
			// 段階 8
			fmt.Println("Brightness: Level 8")
		case val < 0xE666: // 約 52429 - 58981 (9/10)
			// 段階 9
			fmt.Println("Brightness: Level 9")
		default: // 0xE666 から 0xFFFF まで (10/10)
			// 段階 10: 最も明るい、または特定のパターン
			led.Low()
			fmt.Println("Brightness: Level 10 (Highest)")
		}
	}

}
