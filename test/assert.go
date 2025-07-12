package test

import (
	"strings"
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

// Asserts the test string contains the excepted number of substring instances.
func ContainsSubstringCount(t *testing.T, test, substring string, count int, msgAndArgs ...any) {
	assert.Equal(t, count, strings.Count(test, substring), msgAndArgs...)
}
