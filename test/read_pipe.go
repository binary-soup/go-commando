package test

import (
	"os"
)

type PipeReader interface {
	Open(out *os.File)
	Close()
}

type ReadPipe struct {
	reader PipeReader
	in     *os.File
	out    *os.File
}

func NewReadPipe(r PipeReader) ReadPipe {
	in, out, _ := os.Pipe()
	r.Open(out)

	return ReadPipe{
		reader: r,
		in:     in,
		out:    out,
	}
}

func (p ReadPipe) Read(b []byte) (n int, err error) {
	return p.in.Read(b)
}

func (p ReadPipe) Close() {
	p.in.Close()
	p.out.Close()
	p.reader.Close()
}
