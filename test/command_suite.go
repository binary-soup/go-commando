package test

import (
	"github.com/binary-soup/go-command/command"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// CommandSuite is a test suite for commands.
// Provides helpers for various common tests.
type CommandSuite struct {
	suite.Suite
	Cmd command.Command
}

// Creates a new NewCommandSuite for the given command interface.
func NewCommandSuite(cmd command.Command) CommandSuite {
	return CommandSuite{
		Cmd: cmd,
	}
}

// Runs the command with the provided args and requires it passes (returns no error).
func (suite *CommandSuite) RequireCommandPass(args []string) {
	suite.Cmd.SubmitArgs(args)

	err := suite.Cmd.Run()
	require.NoError(suite.T(), err)
}

// Runs the command with the provided args and requires it fails (returns an error).
//
// Also asserts the error message contains all the substrings.
func (suite *CommandSuite) RequireCommandFail(args, errTokens []string) {
	suite.Cmd.SubmitArgs(args)

	err := suite.Cmd.Run()
	require.Error(suite.T(), err)

	AssertErrorContainsSubstrings(suite.T(), err, errTokens)
}
