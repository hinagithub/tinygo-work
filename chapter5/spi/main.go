package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	// spi0 := machine.SPI0 //  汎用SPI
	// spi1 := machine.SPI1 // 液晶ディスプレイ用
	spi := machine.SPI3 // SDカード用
	spi.Configure(machine.SPIConfig{
		SCK:       machine.SPI3_SCK_PIN,
		SDO:       machine.SPI3_SDO_PIN,
		SDI:       machine.SPI3_SDI_PIN,
		Frequency: 4 * 1e6, // 4MHz
	})

	cs := machine.LCD_SS_PIN
	cs.Configure(machine.PinConfig{Mode: machine.PinOutput})
	cs.Low()

	dc := machine.LCD_DC
	dc.Configure(machine.PinConfig{Mode: machine.PinOutput})
	time.Sleep(2 * time.Second) // USB接続待ち

	// Read Display MADCTL
	buf := []byte{0x0B, 0x00}
	dc.Low()             // コマンド
	spi.Tx(buf[:1], nil) // buf[0]を送信する
	dc.High()            // データ
	spi.Tx(nil, buf[1:])
	fmt.Printf("Read Display MADCTL : %X\r\n", buf)

	// Memory Access Control (36h)
	buf = []byte{0x36, 0xAC}
	dc.Low()             // コマンド
	spi.Tx(buf[:1], nil) // buf[0]を送信する
	dc.High()
	spi.Tx(buf[1:], nil) // buf[1]を送信する
	fmt.Printf("Memory Access Control: % X\r\n", buf)

	// Read Display MADCTL (0Bh)
	buf = []byte{0x0B, 0x00}
	dc.Low()             // コマンド
	spi.Tx(buf[:1], nil) // buf[0]を送信する
	dc.High()            // データ
	spi.Tx(nil, buf[:1]) // buf[1]に受信する
	fmt.Printf("Read Display MADCTL: % X\r\n", buf)
	select {}
}
