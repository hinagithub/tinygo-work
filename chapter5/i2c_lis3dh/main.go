package main

import (
	"fmt"
	"machine"

	"tinygo.org/x/drivers/lis3dh"
)

func main() {
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{
		SCL: machine.SCL0_PIN, // クロック用ピン
		SDA: machine.SDA0_PIN, // データ用ピン
	})

	// 初期化
	accel := lis3dh.New(i2c)
	accel.Address = lis3dh.Address0 //Address on wioterminal
	accel.Configure()
	accel.SetRange(lis3dh.RANGE_2_G)

	for {
		// xyz軸の重力速度を読み出し
		x, y, z, _ := accel.ReadAcceleration()
		fmt.Printf("X:%6.2f Y:%6.2f Z:%6.2f\r\n", x, y, z)
	}

	// 出力結果
	// X:%!f(int32= 55677) Y:%!f(int32=-26373) Z:%!f(int32=-991453)
	// X:%!f(int32= 55677) Y:%!f(int32=-26373) Z:%!f(int32=-991453)
	// X:%!f(int32= 55677) Y:%!f(int32=-26373) Z:%!f(int32=-985592)
	// X:%!f(int32= 48840) Y:%!f(int32=-25396) Z:%!f(int32=-985592)
}
