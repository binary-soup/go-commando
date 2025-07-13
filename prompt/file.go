package prompt

import (
	"fmt"
	"os"

	"github.com/binary-soup/go-command/style"
)

// Check if a path already exists, and prompt the user from stdin if they wish to overwrite it.
//
// The prompt will repeat until a valid option is inputted.
// If the path does not exist, no prompt occurs and true is returned.
func (p Prompt) ConfirmOverwrite(title, path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	}

	res := p.ChooseOption(fmt.Sprintf("%s file \"%s\" exists. Overwrite [Y/n]?", style.Bolded.Format(title), path), []byte("Yn"))
	return res == 'Y'
}
