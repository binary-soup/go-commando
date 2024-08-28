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

func (cmd HelloCommand) Run() error {
	fmt.Println("Hello World!")
	return nil
}
