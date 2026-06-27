package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fPath := "./messages.txt"

	file, err := os.Open(fPath)
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
	}

	line := []byte{}
	chunk := make([]byte, 8)

	for {
		n, err := file.Read(chunk)

		if n > 0 {
			for _, b := range chunk[:n] {
				if b == '\n' {
					fmt.Printf("read: %s\n", string(line))
					line = line[:0]
				} else {
					line = append(line, b)
				}
			}
		}

		if err != nil {
			if err == io.EOF {
				if len(line) != 0 {
					fmt.Printf("read: %s\n", string(line))
				}
				return
			}
			fmt.Printf("Error reading the file: %v\n", err)
			return
		}
	}
}
