package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

func main() {
	str := "Hello, world!"
	reader := strings.NewReader(str)
	writer := MyWriter{os.Stdout}

	buffer := make([]byte, 2)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		_, _ = writer.Write(buffer[:n])
	}
}

type MyWriter struct {
	w io.Writer
}

func (mw MyWriter) Write(b []byte) (int, error) {
	return mw.w.Write(b)
}

