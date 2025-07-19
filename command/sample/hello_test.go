package sample_test

import (
	"testing"

	"github.com/binary-soup/go-commando/command/sample"
	"github.com/binary-soup/go-commando/test"
	"github.com/stretchr/testify/suite"
)

type HelloTestSuite struct {
	test.CommandSuite[sample.HelloCommand]
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
	var NAME = test.NewRandFromTime().ASCII(10)

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	s.RequireCommandPass([]string{"-name", NAME})
	test.ContainsSubstrings(s.T(), pipe.NextLine(s.T()), []string{"Hello", NAME})
}
