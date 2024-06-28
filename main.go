package main

import (
	"filetransfer/client"
	"filetransfer/server"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: file-transfer <server|client>")
		return
	}

	switch args[0] {
	case "server":
		server.StartServer()
	case "client":
		if len(args) < 3 {
			fmt.Println("Usage: file-transfer client <server-addr> <file-path>")
			return
		}
		client.StartClient(args[1], args[2])
	default:
		fmt.Println("Invalid argument. Use 'server' or 'client'")
	}
}
