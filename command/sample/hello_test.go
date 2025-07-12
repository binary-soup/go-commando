package sample_test

import (
	"testing"

	"github.com/binary-soup/go-command/command/sample"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/suite"
)

type HelloTestSuite struct {
	test.CommandSuite
}

func TestHelloCommandSuite(t *testing.T) {
	suite.Run(t, &HelloTestSuite{
		CommandSuite: test.NewCommandSuite(sample.NewHelloCommand()),
	})
}

func (s *HelloTestSuite) TestNameNotEmpty() {
	s.RequireCommandFail([]string{}, []string{"name", "cannot", "empty"})
}

func (s *HelloTestSuite) TestPrintName() {
	var NAME = test.RandASCII(s.Rand, 100)

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	s.RequireCommandPass([]string{"-name", NAME})
	test.ContainsSubstrings(s.T(), pipe.NextLine(s.T()), []string{"Hello", NAME})
}
