package util

import (
	"encoding/json"
	"os"
)

// Load the generic type from a JSON file.
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

// Save the generic type to a JSON file.
func SaveJSON[T any](name string, data *T, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return ChainErrorF(err, "error creating %s file", name)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(data)
	if err != nil {
		return ChainErrorF(err, "error encoding %s JSON", name)
	}

	return nil
}
