package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	data := make([]byte, 2048)
	conn.Read(data)
	fmt.Println(string(data))
}
