package util

import (
	"encoding/json"
	"os"
)

func LoadJSON[T any](name, path string) (*T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, ChainErrorF(err, "error opening %s file", name)
	}
	defer file.Close()

	obj := new(T)

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(obj); err != nil {
		return nil, ChainErrorF(err, "error decoding %s JSON", name)
	}

	return obj, nil
}
