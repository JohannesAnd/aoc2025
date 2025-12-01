package utils

import (
	"bufio"
	"os"
)

type LineReader struct {
	scanner *bufio.Scanner
}

func NewLineReader(path string) (*LineReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &LineReader{scanner: bufio.NewScanner(f)}, nil
}

func (lr *LineReader) Next() (string, bool) {
	if lr.scanner.Scan() {
		return lr.scanner.Text(), true
	}
	return "", false
}
