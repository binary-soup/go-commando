// Provides several common helper functions
package alert

import (
	"errors"
	"fmt"

	"github.com/binary-soup/go-commando/style"
)

// Print the error with a styled prefix.
func Print(err error) {
	style.BoldError.Print("ERROR: ")
	fmt.Println(err)
}

// Creates a new error from a string.
func Error(message string) error {
	return errors.New(message)
}

// Creates a new error using fmt package formatting.
func ErrorF(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}

// Chain a new error message with an existing error.
func ChainError(err error, message string) error {
	if err == nil {
		return Error(message)
	}
	return fmt.Errorf("%s\n  %s", message, err.Error())
}

// Chain a new error message using fmt package formatting with an existing error.
func ChainErrorF(err error, format string, args ...any) error {
	return ChainError(err, fmt.Sprintf(format, args...))
}
