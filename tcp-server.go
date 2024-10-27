package main

import (
	"fmt"
	"net"
)

func StartTcpServer(ip string) {
	l, err := net.Listen("tcp", ip+":8080")

	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}

	defer l.Close()

	fmt.Println("TCP server started on", ip+":8080")
	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		fmt.Println("Client connected:", c.RemoteAddr().String())
		go handleTcpConnect(c)
	}
}

func handleTcpConnect(c net.Conn) {
	defer c.Close()

	for {
		var fullMessage []byte
		buffer := make([]byte, 1024)
		for {
			n, err := c.Read(buffer)

			if err != nil {
				fmt.Println("Connection closed by client:", c.RemoteAddr().String())

				return
			}

			fullMessage = append(fullMessage, buffer[:n]...)

			if n < len(buffer) {
				break
			}
		}

		fmt.Println("Received message from client (", len(fullMessage), "bytes)")
		response := fmt.Sprintf("Received %d bytes", len(fullMessage))
		c.Write([]byte(response))
	}
}
