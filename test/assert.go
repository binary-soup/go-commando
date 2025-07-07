package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertContainsAll(t *testing.T, src string, tokens []string) {
	for _, token := range tokens {
		assert.Contains(t, src, token)
	}
}

func AssertErrorContainsAll(t *testing.T, err error, tokens []string) {
	AssertContainsAll(t, err.Error(), tokens)
}
