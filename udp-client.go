package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartUdpClient(ip string) {
	addr, err := net.ResolveUDPAddr("udp", ip+":8080")

	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	c, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}

	defer c.Close()

	buffer := make([]byte, 1024)
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Connected to UDP server:", ip)

	for {
		fmt.Println("Enter message to send to server:")
		text, _ := consoleReader.ReadString('\n')
		text = strings.TrimSpace(text)

		c.Write([]byte(text))

		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			return
		}

		message := string(buffer[:n])
		fmt.Println("Message from server:", message)
	}
}
