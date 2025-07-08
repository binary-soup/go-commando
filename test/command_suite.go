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

func (suite *CommandSuite) AssertCommandSuccess(args []string) {
	err := suite.Cmd.Run(args)
	require.NoError(suite.T(), err)
}

func (suite *CommandSuite) AssertCommandFail(args, errTokens []string) error {
	err := suite.Cmd.Run(args)
	require.Error(suite.T(), err)

	AssertErrorContainsAll(suite.T(), err, errTokens)
	return err
}
