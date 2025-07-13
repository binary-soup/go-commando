package prompt

import (
	"strconv"
)

// Prompt the user from stdin to enter an integer between the range of max and min.
func (p Prompt) Integer(prompt string, min, max int) int {
	for {
		n, err := strconv.Atoi(p.NonEmptyInput(prompt))
		if err != nil || n < min || n > max {
			continue
		}
		return n
	}
}
