package test

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StdoutReader struct {
	stdout  *os.File
	out     *os.File
	in      *os.File
	scanner *bufio.Scanner
}

func NewStdoutReader() StdoutReader {
	r := StdoutReader{
		stdout: os.Stdout,
	}
	r.in, r.out, _ = os.Pipe()

	os.Stdout = r.out
	r.scanner = bufio.NewScanner(r.in)

	return r
}

func (r StdoutReader) ClosePipe() {
	os.Stdout = r.stdout
	r.out.Close()
}

func (r StdoutReader) Close() {
	r.ClosePipe()
	r.in.Close()
}

func (r StdoutReader) NextLine(t *testing.T) string {
	if r.scanner.Scan() {
		return r.scanner.Text()
	}

	assert.FailNow(t, "reached end of stdout")
	return ""
}

func (r StdoutReader) AssertLineContains(t *testing.T, tokens []string) {
	AssertContainsAll(t, r.NextLine(t), tokens)
}
