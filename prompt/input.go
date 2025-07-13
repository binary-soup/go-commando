package prompt

import (
	"fmt"
	"strings"
)

// Prompt the user from stdin to enter any line of text.
func (p Prompt) Input(prompt string) string {
	fmt.Printf("%s ", prompt)
	p.Scanner.Scan()
	return strings.TrimSpace(p.Scanner.Text())
}

// Prompt the user from stdin to enter any non-empty line of text.
//
// The prompt will repeat until a non-empty line is inputted.
func (p Prompt) NonEmptyInput(prompt string) string {
	for {
		line := p.Input(prompt)
		if line != "" {
			return line
		}
	}
}
