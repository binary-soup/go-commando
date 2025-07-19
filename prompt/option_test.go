package prompt_test

import (
	"testing"

	"github.com/binary-soup/go-commando/prompt"
	"github.com/binary-soup/go-commando/test"
	"github.com/stretchr/testify/assert"
)

func TestChooseOption(t *testing.T) {
	var PROMPT = test.NewRandFromTime().ASCII(15)

	in := test.OpenStdinPipe([]any{})
	defer in.Close()

	options := []byte("abcde")
	for i, option := range options {
		// blank, invalid, correct
		var INPUT = []any{"", "X", string(option)}
		in.WriteLines(INPUT)

		pipe := test.OpenStdoutPipe()
		defer pipe.Close()

		res := prompt.New().ChooseOption(PROMPT, options)
		pipe.CloseInput()

		test.PromptCount(t, pipe.NextLine(t), PROMPT, len(INPUT))
		assert.Equal(t, options[i], res, "result does not match chosen option")
	}
}
