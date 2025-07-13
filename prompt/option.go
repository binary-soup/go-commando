package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Prompt the user from stdin to choose from a list of options. Only the first character of input needs to match.
//
// The prompt will repeat until a valid option is selected; then that option is returned.
func ChooseOption(prompt string, options []byte) byte {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s ", prompt)

		scanner.Scan()
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		for _, char := range options {
			if line[0] == char {
				return char
			}
		}
	}
}
