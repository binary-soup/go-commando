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
	suite.AssertCommandFail([]string{}, []string{"name", "cannot", "empty"})
}

func (suite *HelloTestSuite) TestPrintName() {
	const NAME = "Bob"

	r := test.NewStdoutReader()
	defer r.Close()

	suite.AssertCommandSuccess([]string{"-name", NAME})
	r.AssertLineContains(suite.T(), []string{"Hello", NAME})
}
