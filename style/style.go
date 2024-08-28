package style

import (
	"fmt"
	"strings"
)

type Style []string

const (
	Bold      = "1"
	Dim       = "2"
	Underline = "4"

	Black   = "30"
	Red     = "31"
	Green   = "32"
	Yellow  = "33"
	Blue    = "34"
	Magenta = "35"
	Cyan    = "36"
	White   = "37"
)

func New(mod ...string) Style {
	return mod
}

func (sty Style) IsEmpty() bool {
	return len(sty) == 0
}

func (sty Style) Format(s string) string {
	if sty.IsEmpty() {
		return s
	}

	return fmt.Sprintf("\x1B[%sm%s\x1B[0m", strings.Join(sty, ";"), s)
}

func (sty Style) Print(s string) {
	fmt.Print(sty.Format(s))
}

func (sty Style) Println(s string) {
	fmt.Println(sty.Format(s))
}
