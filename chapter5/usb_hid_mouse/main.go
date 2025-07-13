package main

import (
	"machine"
	"machine/usb/hid/mouse"
	"time"
)

func main() {

	button := machine.BUTTON
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	m := mouse.New()

	for {
		if button.Get() {
			m.Press(mouse.Left)
			for i := 0; i < 100; i++ {
				m.Move(1, 0)
				time.Sleep(1 * time.Millisecond)
			}
			for i := 0; i < 100; i++ {
				m.Move(0, 1)
				time.Sleep(1 * time.Millisecond)
			}
			for i := 0; i < 100; i++ {
				m.Move(-1, 0)
				time.Sleep(1 * time.Millisecond)
			}
			for i := 0; i < 100; i++ {
				m.Move(0, -1)
				time.Sleep(1 * time.Millisecond)
			}
			for i := 0; i < 100; i++ {
				m.Move(10, 10)
				time.Sleep(1 * time.Millisecond)
			}

		}
	}

}

func moveMounse() {
	m := mouse.New()
	time.Sleep(5 * time.Second) // パソコンとの接続確率待ち
	m.Move(100, 0)
}
