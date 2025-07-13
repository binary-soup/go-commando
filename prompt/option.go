package prompt

// Prompt the user from stdin to choose from a list of options. Only the first character of input needs to match.
//
// The prompt will repeat until a valid option is selected; then that option is returned.
func (p Prompt) ChooseOption(prompt string, options []byte) byte {
	for {
		line := p.NonEmptyInput(prompt)
		for _, char := range options {
			if line[0] == char {
				return char
			}
		}
	}
}
