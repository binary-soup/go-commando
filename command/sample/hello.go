// Provides sample commands used to demonstrate the command package.
package sample

import (
	"fmt"

	"github.com/binary-soup/go-commando/command"
	"github.com/binary-soup/go-commando/style"
)

// Sample hello command for printing "Hello" to the console.
type HelloCommand struct {
	command.CommandBase
}

// Creates a new HelloCommand.
func NewHelloCommand() HelloCommand {
	return HelloCommand{
		CommandBase: command.NewCommandBase("hello", "prints hello {name} to the console"),
	}
}

// Run the Hello commands. See usage string for details.
func (cmd HelloCommand) Run(args []string) error {
	name := cmd.Flags.String("name", "World", "name to use when saying hello")
	cmd.Flags.Parse(args)

	boldYellow := style.New(style.Bold, style.Yellow)

	fmt.Printf("Hello %s!\n", boldYellow.Format(*name))
	return nil
}
