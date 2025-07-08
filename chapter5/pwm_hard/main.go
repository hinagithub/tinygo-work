// 音階と周波数
//
// | 音名 | 表記 | 周波数     |
// | :--- | :--- | :--------- |
// | ド   | C4   | 261.626 Hz |
// | レ   | D4   | 293.665 Hz |
// | ミ   | E4   | 329.628 Hz |
// | ファ | F4   | 349.228 Hz |
// | ソ   | G4   | 391.995 Hz |
// | ラ   | A4   | 440.000 Hz |
// | シ   | B4   | 493.883 Hz |

package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/tone"
)

// ドからソまで音を変化させる
func main() {
	pwm := machine.TCC0
	pwm.Configure(machine.PWMConfig{})
	channelA, _ := pwm.Channel(machine.BUZZER_CTR)
	notes := []uint64{261, 293, 329, 349, 391}
	i := 0
	for {
		pwm.SetPeriod(1e9 / notes[i])
		pwm.Set(channelA, pwm.Top()/2)
		time.Sleep(100 * time.Millisecond)
		pwm.Set(channelA, 0)
		i = (i + 1) % len(notes)
	}
}

// ラの音を出すだけ
func beepA4() {
	pwm := machine.TCC0
	pwm.Configure(machine.PWMConfig{})
	channelA, _ := pwm.Channel(machine.BUZZER_CTR)
	pwm.SetPeriod(uint64(1e9) / 440) // ラの音
	pwm.Set(channelA, pwm.Top()/2)
	select {}
}

// 定義されたトーンを使った例
func toneExample() {
	speaker, _ := tone.New(machine.TCC0, machine.BUZZER_CTR)
	notes := []tone.Note{tone.A5, tone.B5, tone.C6, tone.D6, tone.F6, tone.G6}
	i := 0
	for {
		speaker.SetNote(notes[i])
		time.Sleep(100 * time.Millisecond)
		i = (i + 1) % len(notes)
	}
}
