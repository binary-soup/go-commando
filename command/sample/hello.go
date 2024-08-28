package sample

import (
	"fmt"
	"local/command"
)

type HelloCommand struct {
	command.CommandBase
}

func NewHelloCommand() HelloCommand {
	return HelloCommand{
		CommandBase: command.NewCommandBase("hello", "prints hello world to the console"),
	}
}

func (cmd HelloCommand) Run(args []string) error {
	name := cmd.Flags.String("name", "World", "name to use when saying hello")
	cmd.Flags.Parse(args)

	fmt.Printf("Hello %s!\n", *name)
	return nil
}
