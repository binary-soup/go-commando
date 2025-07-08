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
	p, s := suite.ConsolePipe()
	defer p.Close()

	const NAME = "Bob"
	suite.AssertCommandSuccess([]string{"-name", NAME})

	s.Scan()
	test.AssertContainsAll(suite.T(), s.Text(), []string{"Hello", NAME})
}
