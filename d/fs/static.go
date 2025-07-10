package fs

import (
	"bytes"
	"io"

	"github.com/binary-soup/go-command/alert"
)

type buffer struct {
	bytes.Buffer
}

func (buffer) Close() error {
	return nil
}

//####################################

type StaticMap map[string]*buffer

func (m StaticMap) Exists(key string) (bool, error) {
	_, ok := m[key]
	return ok, nil
}

func (m StaticMap) Open(key string) (io.ReadCloser, error) {
	buf, ok := m[key]
	if !ok {
		return nil, alert.ErrorF("key \"%s\" not found", key)
	}
	return buf, nil
}

func (m StaticMap) Create(key string) (io.WriteCloser, error) {
	buf := new(buffer)

	m[key] = buf
	return buf, nil
}

func (m StaticMap) Delete(key string) error {
	_, ok := m[key]
	if !ok {
		return alert.ErrorF("key \"%s\" not found", key)
	}

	delete(m, key)
	return nil
}
