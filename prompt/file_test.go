package prompt_test

import (
	"testing"

	"github.com/binary-soup/go-command/prompt"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/assert"
)

func TestConfirmOverwriteFileNotExist(t *testing.T) {
	//var TITLE = test.RandASCII(test.NewRandSource(), 10)
	var PATH = test.TempFile(t, "does/not/exist.txt")

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.ConfirmOverwrite("", PATH)
	pipe.CloseInput()

	pipe.TestEOF(t)
	assert.True(t, res)
}

func TestConfirmOverwriteYes(t *testing.T) {
	var TITLE = test.RandASCII(test.NewRandSource(), 10)
	var PATH = test.CreateEmptyTempFile(t, "file.txt")

	// blank, invalid, wrong case, correct
	in := test.OpenStdinPipe("", "X", "y", "Y")
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.ConfirmOverwrite(TITLE, PATH)

	pipe.CloseInput()
	line := pipe.NextLine(t)

	test.ContainsSubstrings(t, line, []string{TITLE, "exists", PATH})
	test.PromptCount(t, line, TITLE, 4)
	assert.True(t, res)
}

func TestConfirmOverwriteNo(t *testing.T) {
	var TITLE = test.RandASCII(test.NewRandSource(), 10)
	var PATH = test.CreateEmptyTempFile(t, "file.txt")

	// blank, invalid, wrong case, correct
	in := test.OpenStdinPipe("", "X", "N", "n")
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.ConfirmOverwrite(TITLE, PATH)

	pipe.CloseInput()
	line := pipe.NextLine(t)

	test.ContainsSubstrings(t, line, []string{TITLE, "exists", PATH})
	test.PromptCount(t, line, TITLE, 4)
	assert.False(t, res)
}
