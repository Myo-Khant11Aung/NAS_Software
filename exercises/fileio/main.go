package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.WriteFile("test.txt", []byte("Hello NAS"), 0o666)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	}

	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	fmt.Println(string(data))
}
