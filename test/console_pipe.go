package test

import "os"

func NewConsolePipe() ConsolePipe {
	return ConsolePipe{
		ReadPipe:  NewReadPipe(&consoleReader{}),
		WritePipe: NewWritePipe(&consoleWriter{}),
	}
}

type ConsolePipe struct {
	ReadPipe
	WritePipe
}

func (p ConsolePipe) Close() {
	p.ReadPipe.Close()
	p.WritePipe.Close()
}

//####################################

type consoleReader struct {
	stdout *os.File
}

func (r *consoleReader) Open(out *os.File) {
	r.stdout = os.Stdout
	os.Stdout = out
}

func (r consoleReader) Close() {
	os.Stdout = r.stdout
}

//####################################

type consoleWriter struct {
	stdin *os.File
}

func (w *consoleWriter) Open(in *os.File) {
	w.stdin = os.Stdin
	os.Stdin = in
}

func (w consoleWriter) Close() {
	os.Stdin = w.stdin
}
