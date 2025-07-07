package test

import "os"

func CaptureStdIO() StreamCapture {
	return Capture(&stdIOStream{})
}

type stdIOStream struct {
	stdin  *os.File
	stdout *os.File
}

func (s *stdIOStream) Redirect(read, write *os.File) {
	s.stdin = os.Stdin
	s.stdout = os.Stdout

	os.Stdin = read
	os.Stdout = write
}

func (s stdIOStream) Reset() {
	os.Stdin = s.stdin
	os.Stdout = s.stdout
}
