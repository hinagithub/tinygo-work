package main

import (
	"machine/usb/hid/mouse"
	"time"
)

func main() {
	m := mouse.New()
	time.Sleep(1 * time.Second) // パソコンとの接続確率待ち
	m.Move(100, 0)
}
