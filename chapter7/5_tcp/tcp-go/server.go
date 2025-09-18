package main

import (
	"fmt"
	"log"
	"net"
)

var (
	serverIP = "127.0.0.1"
	port     = 8080
)

func main() {
	conn, err := net.Listen(`tcp`, fmt.Sprintf("%s:%d", serverIP, port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("Listen: %s:%d\n", serverIP, port)
	for {
		conn, err := conn.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Printf("from client: %q\r\n", string(buf[:n]))

	// 受信したメッセージを返す
	conn.Write(buf[:n])
	conn.Close()
}
