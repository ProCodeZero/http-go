package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

// Request.Status:
// 1 - initialized
// 2 - done
type Request struct {
	RequestLine RequestLine
	Status      int
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func (r *Request) parse(data []byte) (int, error) {
	r.Status = 1
	return 0, nil
}

func parseRequestLine(data []byte) (requestLine RequestLine, bytesLen int, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == '\r' && data[i+1] == '\n' {
			firstLineList := strings.Split(string(data[:i]), " ")
			fmt.Println(firstLineList)
			httpVer := ""

			if len(firstLineList) != 3 {
				return RequestLine{}, len(data), errors.New("Expected 3 arguments in the Request line")
			}
			for _, l := range firstLineList[0] {
				if !unicode.IsUpper(l) {
					return RequestLine{}, len(data), errors.New("Method should consist of only capital letters")
				}
			}
			for j, l := range firstLineList[2] {
				if l == '/' {
					httpVer = firstLineList[2][j+1:]
					fmt.Println(httpVer)
					if httpVer != "1.1" {
						return RequestLine{}, len(data), errors.New("Http version should be 1.1")
					}
					break
				}
			}

			if httpVer == "" {
				return RequestLine{}, len(data), errors.New("Http version not found")
			}

			return RequestLine{
					Method:        firstLineList[0],
					RequestTarget: firstLineList[1],
					HttpVersion:   httpVer,
				},
				len(data), nil
		}
	}
	// If the new line \r\n wasn't founded
	return RequestLine{}, 0, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return &Request{}, err
	}

	reqLine, bytesLen, err := parseRequestLine(data)
	_ = bytesLen
	if err != nil {
		return &Request{}, err
	}

	return &Request{RequestLine: reqLine}, nil
}
