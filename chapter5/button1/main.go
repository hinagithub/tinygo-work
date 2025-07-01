// 入力端子一覧
// 回路図上の名称  | TinyGoでの定義名   | 端子名  | 初期状態  | 押した状態
// --------------|------------------|--------|----------|----------
// BUTTON1 (右)  | BUTTON_1         | PC26   | High     | Low
// BUTTON2 (中)  | BUTTON_2         | PC27   | High     | Low
// BUTTON3 (左)  | BUTTON_3         | PC28   | High     | Low
// SWITCH_X (上) | SWITCH_X         | PD20   | High     | Low
// SWITCH_Y (左) | SWITCH_Y         | PD12   | High     | Low
// SWITCH_Z (右) | SWITCH_Z         | PD09   | High     | Low
// SWITCH_B (下) | SWITCH_B         | PD08   | High     | Low
// SWITCH_U (押) | SWITCH_U         | PD10   | High     | Low

package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {

	// 端子名で変数化
	// button1 := machine.PC26
	// button2 := machine.PC27
	// button3 := machine.PC28
	// switchX := machine.PD20
	// switchY := machine.PD12
	// switchZ := machine.PD09
	// switchB := machine.PD08
	// switchU := machine.PD10

	// 定義されたラベル名で変数化
	button1 := machine.BUTTON_1
	// button2 := machine.BUTTON_2
	// button3 := machine.BUTTON_3
	// switchX := machine.SWITCH_X
	// switchY := machine.SWITCH_Y
	// switchZ := machine.SWITCH_Z
	// switchB := machine.SWITCH_B
	// switchU := machine.SWITCH_U

	// 初期化
	button1.Configure(machine.PinConfig{Mode: machine.PinInput})
	// button2.Configure(machine.PinConfig{Mode: machine.PinInput})
	// button3.Configure(machine.PinConfig{Mode: machine.PinInput})
	// switchX.Configure(machine.PinConfig{Mode: machine.PinInput})
	// switchY.Configure(machine.PinConfig{Mode: machine.PinInput})
	// switchZ.Configure(machine.PinConfig{Mode: machine.PinInput})
	// switchB.Configure(machine.PinConfig{Mode: machine.PinInput})
	// switchU.Configure(machine.PinConfig{Mode: machine.PinInput})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {

		// ボタンは初期状態でHigh(ture)、押した状態でLow(false)になる「プルアップ入力」
		fmt.Println(button1.Get())
		// ボタンを押すとLEDが消灯する
		led.Set(button1.Get())
		time.Sleep(3 * time.Second)
	}
}
