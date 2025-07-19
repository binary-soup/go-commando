package test

import (
	"github.com/binary-soup/go-commando/command"
	"github.com/stretchr/testify/suite"
)

// CommandSuite is a test suite for commands.
// Provides helpers for running the command with different assertions.
type CommandSuite[T command.Command] struct {
	suite.Suite
	Cmd T
}

// Creates a new NewCommandSuite for the given command interface.
func NewCommandSuite[T command.Command](cmd T) CommandSuite[T] {
	return CommandSuite[T]{
		Cmd: cmd,
	}
}

// Runs the command with the provided args and requires it passes (returns no error).
func (s *CommandSuite[T]) RequireCommandPass(args []string) {
	s.Cmd.SubmitArgs(args)

	err := s.Cmd.Run()
	s.Require().NoError(err)
}

// Runs the command with the provided args and requires it fails (returns an error).
//
// Also asserts the error message contains all the substrings.
func (s *CommandSuite[T]) RequireCommandFail(args, errTokens []string) {
	s.Cmd.SubmitArgs(args)

	err := s.Cmd.Run()
	s.Require().Error(err)

	ErrorContainsSubstrings(s.T(), err, errTokens)
}
