package sample_test

import (
	"testing"

	"github.com/binary-soup/go-command/command/sample"
	"github.com/binary-soup/go-command/test"
	"github.com/stretchr/testify/suite"
)

func TestHelloCommandSuite(t *testing.T) {
	suite.Run(t, &HelloTestSuite{
		CommandSuite: test.NewCommandSuite(sample.NewHelloCommand()),
	})
}

type HelloTestSuite struct {
	test.CommandSuite
}

func (suite *HelloTestSuite) TestNameNotEmpty() {
	suite.RequireCommandFail([]string{}, []string{"name", "cannot", "empty"})
}

func (suite *HelloTestSuite) TestPrintName() {
	const NAME = "Bob"

	out := test.OpenStdoutPipe()
	defer out.Close()

	suite.RequireCommandPass([]string{"-name", NAME})
	out.CloseInput()

	out.AssertLineContainsSubstrings(suite.T(), []string{"Hello", NAME})
}
