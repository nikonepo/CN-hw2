package main

import (
	"fmt"
	"net"
)

func StartUdpServer(ip string) {
	addr, err := net.ResolveUDPAddr("udp", ip+":8080")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	c, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting udp server:", err)
		return
	}

	fmt.Println("UDP server started on", ip+":8080")

	handleUpdConnect(c)
}

func handleUpdConnect(c *net.UDPConn) {
	defer c.Close()

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Println("Received message from client (", len(message), "bytes)")
		response := fmt.Sprintf("Received %d bytes", len(message))
		c.WriteToUDP([]byte(response), clientAddr)
	}
}
