package client

import (
	"fmt"
	"net"
	"os"
)

func StartClient(serverAddr string, filePath string) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	sendFile(conn, file)
}
