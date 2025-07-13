package prompt_test

import (
	"testing"

	"github.com/binary-soup/go-command/prompt"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/suite"
)

type ConfirmOverwriteSuite struct {
	suite.Suite
	Rand test.Rand
}

func TestConfirmOverwriteSuite(t *testing.T) {
	suite.Run(t, &ConfirmOverwriteSuite{
		Rand: test.NewRand(),
	})
}

func (s *ConfirmOverwriteSuite) TestFileNotExist() {
	var PATH = test.TempFile(s.T(), "does/not/exist.txt")

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.New().ConfirmOverwrite("", PATH)
	pipe.CloseInput()

	pipe.TestEOF(s.T())
	s.True(res)
}

func (s *ConfirmOverwriteSuite) TestYes() {
	var TITLE = s.Rand.ASCII(10)
	var PATH = test.CreateEmptyTempFile(s.T(), "file.txt")

	// blank, invalid, wrong case, correct
	var INPUT = []any{"", "X", "y", "Y"}

	in := test.OpenStdinPipe(INPUT...)
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.New().ConfirmOverwrite(TITLE, PATH)

	pipe.CloseInput()
	line := pipe.NextLine(s.T())

	test.ContainsSubstrings(s.T(), line, []string{TITLE, "exists", PATH})
	test.PromptCount(s.T(), line, TITLE, len(INPUT))
	s.True(res)
}

func (s *ConfirmOverwriteSuite) TestNo() {
	var TITLE = s.Rand.ASCII(10)
	var PATH = test.CreateEmptyTempFile(s.T(), "file.txt")

	// blank, invalid, wrong case, correct
	var INPUT = []any{"", "X", "N", "n"}

	in := test.OpenStdinPipe(INPUT...)
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.New().ConfirmOverwrite(TITLE, PATH)

	pipe.CloseInput()
	line := pipe.NextLine(s.T())

	test.ContainsSubstrings(s.T(), line, []string{TITLE, "exists", PATH})
	test.PromptCount(s.T(), line, TITLE, len(INPUT))
	s.False(res)
}
