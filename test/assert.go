package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Asserts the test string contains all the substrings.
func AssertContainsSubstrings(t *testing.T, test string, substrings []string) {
	for _, substring := range substrings {
		assert.Contains(t, test, substring)
	}
}

// Asserts the error message contains all the substrings.
func AssertErrorContainsSubstrings(t *testing.T, err error, substrings []string) {
	AssertContainsSubstrings(t, err.Error(), substrings)
}
