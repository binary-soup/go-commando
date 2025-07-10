package fs

import (
	"io"
	"os"
)

type Disk struct{}

func (Disk) Exists(path string) (bool, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (Disk) Open(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func (Disk) Create(path string) (io.WriteCloser, error) {
	return os.Create(path)
}

func (Disk) Delete(path string) error {
	return os.Remove(path)
}
