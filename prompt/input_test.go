package prompt_test

import (
	"testing"

	"github.com/binary-soup/go-command/prompt"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	src := test.NewRandSource()
	var PROMPT = test.RandASCII(src, 15)
	var INPUT = test.RandASCII(src, 30)

	in := test.OpenStdinPipe(INPUT)
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.New().Input(PROMPT)
	pipe.CloseInput()

	test.PromptCount(t, pipe.NextLine(t), PROMPT, 1)
	assert.Equal(t, INPUT, res)
}

func TestNonEmptyInput(t *testing.T) {
	src := test.NewRandSource()
	var PROMPT = test.RandASCII(src, 15)
	var INPUT = test.RandASCII(src, 30)

	// empty, valid
	in := test.OpenStdinPipe("", INPUT)
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.New().NonEmptyInput(PROMPT)
	pipe.CloseInput()

	test.PromptCount(t, pipe.NextLine(t), PROMPT, 2)
	assert.Equal(t, INPUT, res)
}
