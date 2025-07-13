package prompt

import (
	"bufio"
	"os"
)

type Prompt struct {
	Scanner *bufio.Scanner
}

func New() Prompt {
	return Prompt{
		Scanner: bufio.NewScanner(os.Stdin),
	}
}
