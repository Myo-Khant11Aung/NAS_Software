package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Server couldn't start")
		return
	}
	defer func() {
		if err := listener.Close(); err != nil {
			fmt.Println("Failed to close listener:", err)
		}
	}()

	fmt.Println("Waiting for client...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting client")
			continue
		}

		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("Error closing conn:", err)
		}
	}()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading buffer: ", err)
		return
	}
	fmt.Println(string(buffer[:n]))
	_, err = conn.Write(buffer[:n])
	if err != nil {
		fmt.Println("Error writing data: ", err)
		return
	}
}
