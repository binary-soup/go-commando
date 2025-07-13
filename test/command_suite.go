package test

import (
	"github.com/binary-soup/go-command/command"
	"github.com/stretchr/testify/suite"
)

// CommandSuite is a test suite for commands.
// Provides helpers for various common tests.
type CommandSuite struct {
	suite.Suite
	Rand Rand
	Cmd  command.Command
}

// Creates a new NewCommandSuite for the given command interface.
func NewCommandSuite(cmd command.Command) CommandSuite {
	return CommandSuite{
		Rand: NewRand(),
		Cmd:  cmd,
	}
}

// Runs the command with the provided args and requires it passes (returns no error).
func (s *CommandSuite) RequireCommandPass(args []string) {
	s.Cmd.SubmitArgs(args)

	err := s.Cmd.Run()
	s.Require().NoError(err)
}

// Runs the command with the provided args and requires it fails (returns an error).
//
// Also asserts the error message contains all the substrings.
func (s *CommandSuite) RequireCommandFail(args, errTokens []string) {
	s.Cmd.SubmitArgs(args)

	err := s.Cmd.Run()
	s.Require().Error(err)

	ErrorContainsSubstrings(s.T(), err, errTokens)
}
