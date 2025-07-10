package fs

import "io"

type FileSystem interface {
	Exists(path string) (bool, error)
	Open(path string) (io.ReadCloser, error)
	Create(path string) (io.WriteCloser, error)
	Delete(path string) error
}
