package test

import (
	"fmt"
	"os"
)

// StdinPipe is a test helper for inputting stdin.
// The primary use case is to submit program input typically collected from the console.
type StdinPipe struct {
	stdin *os.File
	out   *os.File
	in    *os.File
}

// Create a new StdinPipe to write input. Takes effect immediately.
//
// input is an optional parameter to write to the pipe after creation.
// After creating the pipe, WriteLines should be used to write further input.
func OpenStdinPipe(input ...any) StdinPipe {
	p := StdinPipe{
		stdin: os.Stdin,
	}
	p.out, p.in, _ = os.Pipe()

	os.Stdin = p.out
	p.WriteLines(input...)

	return p
}

// Close the output and restore stdin.
// The pipe can still be written to, but will no longer be read from.
//
// Has little practical use, but implemented to mirror StdoutPipe.
func (p StdinPipe) CloseOutput() {
	os.Stdin = p.stdin
	p.out.Close()
}

// Close the pipe and restore stdin.
// The pipe can no longer be read from or written to.
func (p StdinPipe) Close() {
	p.CloseOutput()
	p.in.Close()
}

// Write the input as separate lines to the pipe.
func (p StdinPipe) WriteLines(input ...any) {
	for _, line := range input {
		fmt.Fprintln(p.in, line)
	}
}
