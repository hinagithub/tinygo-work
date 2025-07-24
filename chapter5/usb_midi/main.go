// USB MIDI = Unibersal Serial Bus Musical Instrument Digital Interface
// 演奏情報をUSBでやりとりする
// MIDIはMIDIが「音そのもの」ではなく、「音を出すための指示や情報」を伝達する規格である

// MIDIが伝える「情報」の例
// ノートオン（Note On）: どの鍵盤が（音程）、どのくらいの強さで（ベロシティ）、弾かれたか
// 例：「中央のCの音を、強く弾き始めた」
// ノートオフ（Note Off）: どの鍵盤が、いつ離されたか
// 例：「中央のCの音を、離した」
// プログラムチェンジ（Program Change）: 楽器の音色を何番に変えるか
// 例：「ピアノの音色から、ストリングスの音色に変える」
// コントロールチェンジ（Control Change）: 音量、パン（左右の定位）、ビブラートの深さなど、様々な音のパラメーターをどう変化させるか
// 例：「音量を少しずつ上げる」「サスティンペダルを踏んだ」
// ピッチベンド（Pitch Bend）: 音程を滑らかに上下させる（ギターのチョーキングのように）

// MIDI Keyboard
// https://www.onlinemusictools.com/kb/

package main

import (
	"machine"
	"machine/usb/adc/midi"
	"time"
)

func main() {
	led := machine.LCD_BACKLIGHT
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m := midi.New()
	m.SetHandler(func(b []byte) {
		// 外部からのMIDI信号を受信
		led.Toggle()
	})
	time.Sleep(1 * time.Second)

	for {
		// 外部へMIDI信号を送信
		m.NoteOn(0, 0, midi.C4, 0x40)
		time.Sleep(time.Millisecond * 1000)
		m.NoteOff(0, 0, midi.C4, 0x40)
		time.Sleep(time.Millisecond * 1000)
	}
}
