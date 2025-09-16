package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, port))
	if err != nil {
		log.Fatal(err)
	}

	// サーバに送信
	fmt.Fprintf(conn, "hello from go\r\n")

	// サーバからの受信を標準出力に書き出す
	io.Copy(os.Stdout, conn)
	conn.Close()
}
