// I2C = Inter Integrated Circuit 
// ATSUMD51ではSERCOMというペリフェラル名

package main

import (
	"machine"
	"time"
	"fmt"
)

func main() {
	// I2Cの初期化
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ, // 400kHz
	})

	// デバイスのアドレス
	const DEVICE_ADDR = 0x44 // 例: SHT30温湿度センサー

	// データの送信例
	txData := []byte{0x2C, 0x06} // コマンド
	err := machine.I2C0.Tx(DEVICE_ADDR, txData, nil)
	if err != nil {
		fmt.Println("送信エラー:", err)
		return
	}

	time.Sleep(15 * time.Millisecond) // センサーの処理待ち

	// データの受信例
	rxData := make([]byte, 6) // 6バイト受信
	err = machine.I2C0.Tx(DEVICE_ADDR, nil, rxData)
	if err != nil {
		fmt.Println("受信エラー:", err)
		return
	}

	// 受信データの表示
	fmt.Printf("受信データ: %x\n", rxData)

	// I2Cバスのスキャン例
	scanI2CBus()
}

// I2Cバス上のデバイスをスキャンする関数
func scanI2CBus() {
	fmt.Println("I2Cデバイススキャン開始...")
	
	found := 0
	for addr := uint8(0x08); addr < 0x78; addr++ {
		// 各アドレスにテスト送信
		err := machine.I2C0.Tx(uint16(addr), []byte{}, nil)
		if err == nil {
			fmt.Printf("デバイス発見: 0x%02X\n", addr)
			found++
		}
	}
	
	if found == 0 {
		fmt.Println("デバイスが見つかりませんでした")
	} else {
		fmt.Printf("合計 %d 個のデバイスが見つかりました\n", found)
	}
}
