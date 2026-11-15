package main

import (
	"bufio"
	"fmt"
	"io"
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
			fmt.Println("Error accepting client: ", err)
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

	reader := bufio.NewReader(conn)

	for {
		incomingMessage, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected gracefully")
				return
			} else {
				fmt.Println("Reading Error: ", err)
				return
			}
		}

		_, err = conn.Write([]byte(incomingMessage))
		if err != nil {
			fmt.Println("Writing err: ", err)
			return
		}
	}
}
