package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
