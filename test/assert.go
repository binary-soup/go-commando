package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Asserts the test string contains all the substrings.
func ContainsSubstrings(t *testing.T, test string, substrings []string) {
	for _, substring := range substrings {
		assert.Contains(t, test, substring)
	}
}

// Asserts the error message contains all the substrings.
func ErrorContainsSubstrings(t *testing.T, err error, substrings []string) {
	ContainsSubstrings(t, err.Error(), substrings)
}
