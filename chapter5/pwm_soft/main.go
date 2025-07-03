// PWM = Plus Width Modulation

package main

import (
	"machine"
	"time"
)

func main() {
	bzr := machine.BUZZER_CTR
	bzr.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		bzr.Toggle()
		// ↑高音
		// time.Sleep(100 * time.Microsecond) // 5000Hz
		// time.Sleep(1 * time.Millisecond) // 500Hz
		// time.Sleep(10 * time.Millisecond) // 50Hz
		time.Sleep(100 * time.Millisecond) // 5Hz
		// ↓低音
	}
}
