// USB_HID
// Universal Serial Bus Human Interface Device
// 人間のインターフェースデバイス = キーボードやマウス
// wioターミナルをキーボードとして認識する

package main

import (
	"machine"
	"machine/usb/hid/keyboard"
	"time"
)

func main() {

	button := machine.BUTTON
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	kb := keyboard.New()

	for {
		// button.Get() ・・・ ボタンを押してればtrue、押していなければfalse
		if !button.Get() {
			kb.Write([]byte("echo "))

			// Tiny
			kb.Down(keyboard.KeyModifierShift) // Shiftキー押下
			kb.Press(keyboard.KeyT)
			kb.Up(keyboard.KeyModifierShift) // Shiftキーを離す
			kb.Press(keyboard.KeyI)
			kb.Press(keyboard.KeyN)
			kb.Press(keyboard.KeyY)

			// Go
			kb.Down(keyboard.KeyModifierShift) // Shiftキー押下
			kb.Press(keyboard.KeyG)
			kb.Up(keyboard.KeyModifierShift) // Shiftキーを離す
			kb.Press(keyboard.KeyO)

			kb.Press(keyboard.KeyReturn)
			time.Sleep(200 * time.Millisecond)

		}
	}
}

// キーボードでTinyGoと入力するだけ
func writeTinyGo() {
	kb := keyboard.New()
	time.Sleep(1 * time.Second)
	kb.Write([]byte("TynyGo"))
}
