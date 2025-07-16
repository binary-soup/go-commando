package test

import (
	"testing"

	"github.com/binary-soup/go-command/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Load a json object from the path and assert it matches the expected object.
func CompareJSONFile[T any](t *testing.T, path string, obj T) {
	require.FileExists(t, path, "JSON file does not exist")

	obj2, err := data.LoadJSON[T]("test json", path)
	require.NoError(t, err)

	assert.Equal(t, obj, obj2, "JSON objects are not the same")
}
