package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var target string
	// TODO: Add option to supply or link file with IP's
	fmt.Print("Enter IP address: ")
	fmt.Scanln(&target)

	startTime := time.Now()

	const startPort = 1
	const endPort = 1024
	totalPorts := endPort - startPort + 1

	fmt.Printf("\nScanning %s...\n\n", target)

	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf("%s:%d", target, port)

		conn, err := net.DialTimeout(
			"tcp",
			address,
			200*time.Millisecond,
		)

		if err == nil {
			fmt.Printf("\nPort %d is OPEN\n", port)
			conn.Close()
		}

		// Calculaterr progress
		scanned := port - startPort + 1
		percentage := float64(scanned) / float64(totalPorts) * 100

		// \r moves the cursor back to the beginning of the line fixes breaking progress bar
		fmt.Printf("\rScanning port %d/%d (%.1f%%)", port, endPort, percentage)
	}

	elapsed := time.Since(startTime)

	fmt.Printf("\n\nScan complete.\n")
	fmt.Printf("Time taken: %v\n", elapsed)
}
