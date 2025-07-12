package prompt_test

import (
	"testing"

	"github.com/binary-soup/go-command/prompt"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/assert"
)

func TestPromptOption(t *testing.T) {
	var PROMPT = test.RandASCII(test.NewRandSource(), 15)

	in := test.OpenStdinPipe()
	defer in.Close()

	options := []byte("abcde")
	for i, option := range options {
		in.WriteLines("", "X", string(option))

		pipe := test.OpenStdoutPipe()
		defer pipe.Close()

		res := prompt.ChooseOption(PROMPT, options)
		pipe.CloseInput()

		test.ContainsSubstringCount(t, pipe.NextLine(t), PROMPT, 3, "wrong number of prompts")
		assert.Equal(t, options[i], res, "result does not match chosen option")
	}
}
