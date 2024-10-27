package main

import (
	"flag"
	"fmt"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP address")
	mode := flag.String("mode", "tcp-server", "Mode: tcp-server, tcp-client, udp-server, udp-client")
	flag.Parse()

	switch *mode {
	case "tcp-server":
		StartTcpServer(*ip)
	case "tcp-client":
		StartTcpClient(*ip)
	case "udp-server":
		StartUdpServer(*ip)
	case "udp-client":
		StartUdpClient(*ip)
	default:
		fmt.Println("Illegal mode")
	}
}
