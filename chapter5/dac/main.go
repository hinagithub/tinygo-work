// DACピンマッピング情報
// -----|--------|------------------------
// DAC  | ピン名  |  Wio Terminalでの端子名
// -----|--------|------------------------
// DAC0 | PA02   | BCM17
// DAC1 | PA05   | BCM7
// -----|--------|------------------------

package main

import (
	"machine"
	"time"
)

func main() {
	dac := machine.DAC0
	dac.Configure(machine.DACConfig{})
	for {
		dac.Set(0xFFF0) // 約3.3V
		time.Sleep(100 * time.Millisecond)
		dac.Set(0x8000) // 約1.6V
		time.Sleep(100 * time.Millisecond)
		dac.Set(0x4000) // 約0.8V
		time.Sleep(100 * time.Millisecond)
	}
}
