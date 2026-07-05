package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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
	fPath := "./messages.txt"

	file, err := os.Open(fPath)
	if err != nil {
		log.Fatalf("Error opening the file: %v\n", err)
	}

	lines := getLinesChannel(file)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
