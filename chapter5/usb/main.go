package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	time.Sleep(5 * time.Second)
	fmt.Printf("hello tinygo\r\n")

	msg := ""
	fmt.Scanf("msg %q\r\n", &msg)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("You typed: %s\r\n", scanner.Text())
	}
}
