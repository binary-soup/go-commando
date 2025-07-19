package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/binary-soup/go-commando/data"
	"github.com/stretchr/testify/require"
)

// Create a new temp directory and join with path.
func TempFile(t *testing.T, path string) string {
	return filepath.Join(t.TempDir(), path)
}

// Create the file in a new temp directory.
// Returns the open file pointer and the path.
func CreateTempFile(t *testing.T, path string) (*os.File, string) {
	path = TempFile(t, path)

	file, err := os.Create(path)
	require.NoError(t, err)

	return file, path
}

// Create an empty file in a new temp directory and return its path.
//
// Note: the file is closed after creation. If you wish to write to the new file, use CreateTempFile instead.
func CreateEmptyTempFile(t *testing.T, path string) string {
	file, path := CreateTempFile(t, path)
	file.Close()

	return path
}

// Save the json object to a new temp directory and return its path.
func CreateJSONTempFile[T any](t *testing.T, path string, obj T) string {
	path = TempFile(t, path)

	err := data.SaveJSON("test json", obj, path)
	require.NoError(t, err)

	return path
}
