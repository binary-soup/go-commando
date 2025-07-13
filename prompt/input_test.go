package prompt_test

import (
	"testing"

	"github.com/binary-soup/go-command/prompt"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/suite"
)

type InputSuite struct {
	suite.Suite
	Rand test.Rand
}

func TestInputSuite(t *testing.T) {
	suite.Run(t, &InputSuite{
		Rand: test.NewRand(),
	})
}

func (s *InputSuite) TestInput() {
	var PROMPT = s.Rand.ASCII(15)
	var TEXT = s.Rand.ASCII(30)

	in := test.OpenStdinPipe(TEXT)
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.New().Input(PROMPT)
	pipe.CloseInput()

	test.PromptCount(s.T(), pipe.NextLine(s.T()), PROMPT, 1)
	s.Equal(TEXT, res)
}

func (s *InputSuite) TestNonEmptyInput() {
	var PROMPT = s.Rand.ASCII(15)
	var TEXT = s.Rand.ASCII(30)

	// empty, valid
	var INPUT = []any{"", TEXT}

	in := test.OpenStdinPipe(INPUT...)
	defer in.Close()

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	res := prompt.New().NonEmptyInput(PROMPT)
	pipe.CloseInput()

	test.PromptCount(s.T(), pipe.NextLine(s.T()), PROMPT, len(INPUT))
	s.Equal(TEXT, res)
}
