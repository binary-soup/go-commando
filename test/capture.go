package test

import (
	"os"
)

type Stream interface {
	Redirect(read, write *os.File)
	Reset()
}

type StreamCapture struct {
	stream Stream
	read   *os.File
	write  *os.File
}

func Capture(s Stream) StreamCapture {
	r, w, _ := os.Pipe()

	c := StreamCapture{
		stream: s,
		read:   r,
		write:  w,
	}

	s.Redirect(r, w)
	return c
}

func (s StreamCapture) Read(p []byte) (n int, err error) {
	return s.read.Read(p)
}

func (s StreamCapture) Close() {
	s.write.Close()
	s.stream.Reset()
}
