package sample

import (
	"fmt"
	"local/command"
	"local/style"
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

	boldYellow := style.New(style.Bold, style.Yellow)

	fmt.Printf("Hello %s!\n", boldYellow.Format(*name))
	return nil
}
