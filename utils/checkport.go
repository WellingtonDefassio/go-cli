package utils

import (
	"fmt"
	"net"
	"time"
)

func CheckPort(host string, ports []string) {
	for _, port := range ports {
		second := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), second)
		if err != nil {
			fmt.Printf("port %s is closed\n", port)
		}
		if conn != nil {
			fmt.Printf("port %s is open\n", port)
			defer conn.Close()
		}
	}
}
