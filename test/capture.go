package test

import (
	"os"
)

type Stream interface {
	Redirect(w *os.File)
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

	s.Redirect(w)
	return c
}

func (s StreamCapture) Read(p []byte) (n int, err error) {
	return s.read.Read(p)
}

func (s StreamCapture) Close() {
	s.write.Close()
	s.stream.Reset()
}

//############################

type stdoutStream struct {
	stdout *os.File
}

func (s *stdoutStream) Redirect(w *os.File) {
	s.stdout = os.Stdout
	os.Stdout = w
}

func (s stdoutStream) Reset() {
	os.Stdout = s.stdout
}

func CaptureStdout() StreamCapture {
	return Capture(&stdoutStream{})
}
