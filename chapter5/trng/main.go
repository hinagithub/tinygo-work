// TRNG = True Random Number Generator
package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	var result [4]byte
	for {
		rand.Read(result[:])
		fmt.Printf("% X\n\r", result)
		time.Sleep(time.Second)
	}
}
