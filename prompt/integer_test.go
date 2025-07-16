package prompt_test

import (
	"testing"

	"github.com/binary-soup/go-command/prompt"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/assert"
)

func TestInteger(t *testing.T) {
	r := test.NewRandFromTime()

	var PROMPT = r.ASCII(15)
	var LIMIT = r.IntRange(-1000, 1000)

	// empty, not an int, too small, too big, valid
	var INPUT = []any{"", "X", LIMIT - 1, LIMIT + 1, LIMIT}

	in := test.OpenStdinPipe(INPUT)
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.New().Integer(PROMPT, LIMIT, LIMIT)
	pipe.CloseInput()

	test.PromptCount(t, pipe.NextLine(t), PROMPT, len(INPUT))
	assert.Equal(t, LIMIT, res)
}
