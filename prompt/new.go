package prompt

import (
	"bufio"
	"os"
)

// Provides several helper methods for prompting the user from stdin.
type Prompt struct {
	Scanner *bufio.Scanner
}

// Create a new Prompt object.
func New() Prompt {
	return Prompt{
		Scanner: bufio.NewScanner(os.Stdin),
	}
}
