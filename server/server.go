package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Waiting for connection...")
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
	conn.Write([]byte("Hello, Peer!"))
}
