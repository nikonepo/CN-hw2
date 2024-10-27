package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartTcpClient(ip string) {
	c, err := net.Dial("tcp", ip+":8080")

	if err != nil {
		fmt.Println("Error connecting to TCP server:", err)
		return
	}

	defer c.Close()

	buffer := make([]byte, 1024)
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Connected to TCP server:", ip)

	for {
		fmt.Println("Enter message to send to server:")
		text, _ := consoleReader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "big" {
			largeMessage := strings.Repeat("A", 40000)
			c.Write([]byte(largeMessage))
			fmt.Println("Sent large message to server")
		} else {
			c.Write([]byte(text))
		}

		n, err := c.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from TCP connection:", err)
			return
		}

		message := string(buffer[:n])
		fmt.Println("Message from server:", message)
	}
}
