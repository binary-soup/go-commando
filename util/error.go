// Provides several common helper functions
package util

import (
	"errors"
	"fmt"
)

// Creates a new error from a string.
func Error(message string) error {
	return errors.New(message)
}

// Chain a new error message with an existing error.
func ChainError(err error, message string) error {
	return fmt.Errorf("%s\n  %s", message, err.Error())
}

// Chain a new error message using fmt package formatting with an existing error.
func ChainErrorF(err error, format string, args ...any) error {
	return ChainError(err, fmt.Sprintf(format, args...))
}
