package main

import (
	"fmt"
	"time"

	"github.com/sago35/tinygo-examples/tree/main/wioterminal/initialize"
)

func main() {
	updateCheck()
}

func updateCheck() {
	rtl, err := initialize.SetupRTL8720DN
	if err != nil {
		failMessage(err.Error())
	}

	ver, err := rtl.Version()
	if err != nil {
		failMessage(err.Error())
	}
	for {
		fmt.Printf("RTL8270DN Firware Version %s\r\n", ver)
		time.Sleep(10 * time.Second)
	}

}

func failMessage(str string) {
	for {
		fmt.Println("%s\n", str)
		time.Sleep(5 * time.Second)
	}
}
