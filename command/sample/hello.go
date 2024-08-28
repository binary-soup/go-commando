// Provides sample commands used to demonstrate the command package.
package sample

import (
	"fmt"

	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/style"
)

// Sample hello command for printing "Hello" to the console.
type HelloCommand struct {
	command.CommandBase
}

// Creates a new HelloCommand.
func NewHelloCommand() HelloCommand {
	return HelloCommand{
		CommandBase: command.NewCommandBase("hello", "prints hello world to the console"),
	}
}

// Prints "Hello {name}!" to the console.
// Args: -name (default "World").
func (cmd HelloCommand) Run(args []string) error {
	name := cmd.Flags.String("name", "World", "name to use when saying hello")
	cmd.Flags.Parse(args)

	boldYellow := style.New(style.Bold, style.Yellow)

	fmt.Printf("Hello %s!\n", boldYellow.Format(*name))
	return nil
}
