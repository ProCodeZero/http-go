package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		line := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}

			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				line += string(data[:i])
				out <- line

				line = ""
				line += string(data[i+1:])
				continue
			}

			line += string(data)

		}
		if len(line) != 0 {
			out <- line
		}
	}()

	return out
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("Error listening tcp: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting conn: %v", err)
			continue
		}

		fmt.Println("Connection accepted")

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Println(line)
		}

		fmt.Println("Connection has been closed")
	}
}
