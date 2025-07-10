package data

import (
	"encoding/json"

	"github.com/binary-soup/go-command/alert"
	"github.com/binary-soup/go-command/d"
)

// Load the generic type from a JSON file.
func LoadJSON[T any](name, path string) (T, error) {
	var obj T

	file, err := d.FileSystem.Open(path)
	if err != nil {
		return obj, alert.ChainErrorF(err, "error opening %s file", name)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&obj); err != nil {
		return obj, alert.ChainErrorF(err, "error decoding %s JSON", name)
	}

	return obj, nil
}

// Save the generic type to a JSON file.
func SaveJSON[T any](name string, obj T, path string) error {
	file, err := d.FileSystem.Create(path)
	if err != nil {
		return alert.ChainErrorF(err, "error creating %s file", name)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(obj)
	if err != nil {
		return alert.ChainErrorF(err, "error encoding %s JSON", name)
	}

	return nil
}
