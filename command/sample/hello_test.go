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

func (suite *HelloTestSuite) TestNameNotEmpty() {
	suite.RequireCommandFail([]string{}, []string{"name", "cannot", "empty"})
}

func (suite *HelloTestSuite) TestPrintName() {
	var NAME = test.RandASCII(suite.Rand, 100)

	pipe := test.OpenStdoutPipe()
	defer pipe.Close()

	suite.RequireCommandPass([]string{"-name", NAME})
	pipe.TestNextLineContainsSubstrings(suite.T(), []string{"Hello", NAME})
}
