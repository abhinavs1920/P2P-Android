package server

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

	// Get the IP address of the server
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return
	}

	// Display all IP addresses for the server
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			fmt.Printf("Server started at: http://%s:8080\n", ipNet.IP.String())
		}
	}

	http.Serve(listener, nil)
}
