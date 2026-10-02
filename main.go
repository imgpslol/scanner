package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var target string

	fmt.Print("Enter IP address: ")
	fmt.Scanln(&target)

	for port := 1; port <= 1024; port++ {
		address := fmt.Sprintf("%s:%d", target, port)

		conn, err := net.DialTimeout(
			"tcp",
			address,
			200*time.Millisecond,
		)

		if err == nil {
			fmt.Printf("Port %d is OPEN\n", port)
			conn.Close()
		}
	}
}
