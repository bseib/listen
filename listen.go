package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	const ADDR = "localhost"
	const PORT = "62869"
	listener, err := net.Listen("tcp", ADDR + ":" + PORT)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	defer listener.Close()
	fmt.Println("Listening on " + ADDR + ":" + PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(3)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		bytes, err := conn.Read(buf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Stdout.Write(buf[0:bytes])
	}
}
