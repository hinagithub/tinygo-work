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
	"machine"
	"time"
)

func main() {

	button1 := machine.BUTTON_1
	button1.Configure(machine.PinConfig{Mode: machine.PinInput})
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button1.SetInterrupt(machine.PinToggle, func(machine.Pin) {
		// button1が変化した時の実装を書く
		led.Set(button1.Get())
	})

	for {
		time.Sleep(1 * time.Second)
	}
}

// 割り込みとgoroutineはうまく使い分ける必要あり
// 割り込み: イベント発火後すぐに呼ばれる
// goroutine: runtime.Gosched()やtime.Sleep()が呼ばれた後に処理される
