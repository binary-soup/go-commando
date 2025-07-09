package test

import (
	"github.com/binary-soup/go-command/command"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CommandSuite struct {
	suite.Suite
	Cmd command.Command
}

func NewCommandSuite(cmd command.Command) CommandSuite {
	return CommandSuite{
		Cmd: cmd,
	}
}

func (suite *CommandSuite) RequireCommandSuccess(args []string) {
	suite.Cmd.SubmitArgs(args)

	err := suite.Cmd.Run()
	require.NoError(suite.T(), err)
}

func (suite *CommandSuite) RequireCommandFail(args, errTokens []string) {
	suite.Cmd.SubmitArgs(args)

	err := suite.Cmd.Run()
	require.Error(suite.T(), err)

	AssertErrorContainsAll(suite.T(), err, errTokens)
}
