package test

import (
	"testing"

	"github.com/binary-soup/go-command/config"
	"github.com/stretchr/testify/require"
)

// Assert the config's validate methods returns no validation errors.
func ConfigValid(t *testing.T, cfg config.Config) {
	ConfigValidateErrors(t, cfg)
}

// Assert the config's validate methods returns the expected number of validation errors;
// and each error contains its respective substrings.
func ConfigValidateErrors(t *testing.T, cfg config.Config, errSubstrings ...[]string) {
	verrs := cfg.Validate()

	require.Equal(t, len(verrs), len(errSubstrings))
	for i, substrings := range errSubstrings {
		ContainsSubstrings(t, verrs[i].Error(), substrings)
	}
}
