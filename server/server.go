package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

var transferHistory []TransferRecord
var historyMutex sync.Mutex

func StartServer() {
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/history", handleHistory)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server started at:", listener.Addr().String())
	http.Serve(listener, nil)
}
