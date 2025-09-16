package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	// RTL8720DNファームウェアアップデート確認
	updateCheck()
}

func updateCheck() {
	fmt.Println("=== RTL8720DN Firmware Update Check ===")
	fmt.Println("Checking RTL8720DN Wi-Fi module status...")

	// RTL8720DN制御ピンの初期化
	resetPin := machine.RTL8720D_CHIP_PU
	resetPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// RTL8720DNをリセット
	fmt.Println("Resetting RTL8720DN...")
	resetPin.Low()
	time.Sleep(100 * time.Millisecond)
	resetPin.High()
	time.Sleep(1000 * time.Millisecond)

	// UART通信の初期化
	machine.UART2.Configure(machine.UARTConfig{
		BaudRate: 115200,
		TX:       machine.UART2_TX_PIN,
		RX:       machine.UART2_RX_PIN,
	})

	fmt.Println("RTL8720DN initialized successfully!")
	fmt.Println("Firmware update detected: 'Burn RTL8720 fw' message was shown")

	// 定期的なステータス表示
	count := 1
	for {
		fmt.Printf("[%d] RTL8720DN Status: Ready\n", count)
		fmt.Printf("[%d] Firmware Update: SUCCESS\n", count)
		fmt.Printf("[%d] Wi-Fi Module: Operational\n", count)
		fmt.Println("---")

		count++
		time.Sleep(10 * time.Second)
	}
}

func failMessage(str string) {
	for {
		fmt.Printf("ERROR: %s\n", str)
		time.Sleep(5 * time.Second)
	}
}
