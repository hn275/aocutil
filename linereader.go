package aocutil

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type LineReader struct {
	fd  *os.File
	buf bufio.Reader
}

func NewLineReader(filename string) *LineReader {
	var err error
	lr := new(LineReader)
	lr.fd, err = os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file %s\n%s\n", filename, err.Error())
		os.Exit(1)
	}
	lr.buf = *bufio.NewReader(lr.fd)
	return lr
}

// return null at end of file
func (lr *LineReader) NextLine() []byte {
	line, _, err := lr.buf.ReadLine()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		fmt.Fprintf(os.Stderr, "failed bufio.ReadLine %s\n", err.Error())
		os.Exit(1)

	}
	return line
}
