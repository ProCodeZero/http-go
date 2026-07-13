package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		log.Fatalf("Error resolving UDP addr: %v\n", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Printf("Error dial udp: %v", err)
	}
	defer conn.Close()

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		str, err := r.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v", err)
		}
		_, err = conn.Write([]byte(str))
		if err != nil {
			log.Printf("Error sending input: %v", err)
		}
	}
}
