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

	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{
		SCL: machine.SCL0_PIN, // クロック用ピン
		SDA: machine.SDA0_PIN, // データ用ピン
	})

	// 初期化
	i2c.WriteRegister(0x18, 0x20, []byte{0x57})
	data := []byte{0, 0, 0, 0, 0, 0}
	for {
		// XYZ軸の重力加速度を読み出し
		i2c.ReadRegister(0x18, 0x28|0x80, data)

		x := readAcceleration(data[0], data[1])
		y := readAcceleration(data[2], data[3])
		z := readAcceleration(data[4], data[5])

		fmt.Printf("X:%6.2f Y:%6.2f Z:%6.2f", x, y, z)

		time.Sleep(100 * time.Millisecond)
	}
}

func readAcceleration(l, h byte) float32 {
	// uint8の値を組み合わせてuint16の変数を作成
	a := uint16(l) | uint16(h)<<8
	
	// 符号付き16bitとして解釈（2の補数）
	signed := int16(a)
	
	//0x4000==1gのため0x4000で割る
	return float32(signed) / 0x4000

	// 例: 
	// l = 0x34
	// h = 0x12

	// uint16(l)        // 0000 0000 0011 0100 (0x0034)
	// uint16(h)<<8     // 0001 0010 0000 0000 (0x1200)

	// ビット単位のOR演算
	//   0000 0000 0011 0100  (0x0034)
	// | 0001 0010 0000 0000  (0x1200)
	// ─────────────────────
	//   0001 0010 0011 0100  (0x1234) ← これがaに入る
}

// 接続するだけ
func connect() {
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{
		SCL: machine.SCL0_PIN, // クロック用ピン
		SDA: machine.SDA0_PIN, // データ用ピン
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
