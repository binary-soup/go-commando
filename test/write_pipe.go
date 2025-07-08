package test

import (
	"os"
)

type PipeWriter interface {
	Open(in *os.File)
	Close()
}

type WritePipe struct {
	writer PipeWriter
	in     *os.File
	out    *os.File
}

func NewWritePipe(w PipeWriter) WritePipe {
	in, out, _ := os.Pipe()
	w.Open(in)

	return WritePipe{
		writer: w,
		in:     in,
		out:    out,
	}
}

func (p WritePipe) Write(b []byte) (n int, err error) {
	return p.out.Write(b)
}

func (p WritePipe) Close() {
	p.in.Close()
	p.out.Close()
	p.writer.Close()
}
