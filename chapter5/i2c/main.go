// I2C = Inter Integrated Circuit
// ATSUMD51ではSERCOMというペリフェラル名

// Example of acceleration data
// Table 9: Output data registers content vs. acceleration (FS = ±2 g, high-resolution mode)
//
// ┌─────────────────┬─────────────────────────────────────────────────────────────┐
// │                 │                     Register address                        │
// │ Acceleration    ├─────────────────────────┬───────────────────────────────────┤
// │ values          │        BLE = 0          │            BLE = 1                │
// │                 ├─────────┬───────────────┼─────────┬─────────────────────────┤
// │                 │   28h   │      29h      │   28h   │         29h             │
// ├─────────────────┼─────────┼───────────────┼─────────┼─────────────────────────┤
// │      0 g        │   00h   │      00h      │   00h   │         00h             │
// │    350 mg       │   E0h   │      15h      │   15h   │         E0h             │
// │      1 g        │   00h   │      40h      │   40h   │         00h             │
// │   -350 mg       │   20h   │      EAh      │   EAh   │         20h             │
// │     -1 g        │   00h   │      C0h      │   C0h   │         00h             │
// └─────────────────┴─────────┴───────────────┴─────────┴─────────────────────────┘

package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {

}

// 接続するだけ
func connect() {
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{
		SCL: machine.SCL0_PIN,
		SDA: machine.SDA0_PIN,
	})
	time.Sleep(2 * time.Second) // USB CDC接続待ち
	// 読み出し用のスライスを用意して読み込む
	data := []byte{0}
	err := i2c.ReadRegister(0x18, 0x0F, data)
	if err != nil {
		// エラー処理を行う
	}
	fmt.Printf("Who am I : 0x%02x\r\n", data[0])
	// Who am I : 0x33
	select {}
}
