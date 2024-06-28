package client

import (
	"net"
	"time"
)

func discoverServer() (string, error) {
	conn, err := net.DialTimeout("udp", "255.255.255.255:8081", time.Second*2)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte("DISCOVER"))
	if err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}
