// Provides sample commands used to demonstrate the command package.
package sample

import (
	"fmt"

	"github.com/binary-soup/go-command/alert"
	"github.com/binary-soup/go-command/command"
	"github.com/binary-soup/go-command/style"
)

// Sample hello command for printing "Hello" to the console.
type HelloCommand struct {
	command.CommandBase
	name *string
}

// Creates a new HelloCommand.
func NewHelloCommand() HelloCommand {
	base := command.NewCommandBase("hello", "prints hello {name} to the console")

	return HelloCommand{
		CommandBase: base,
		name:        base.Flags.String("name", "", "name to use when saying hello"),
	}
}

// Run the Hello commands. See usage string for details.
func (cmd HelloCommand) Run(args []string) error {
	cmd.ParseFlags(args)

	if *cmd.name == "" {
		return alert.Error("\"name\" cannot be empty")
	}

	fmt.Printf("Hello %s!\n", style.New(style.Bold, style.Yellow).Format(*cmd.name))
	return nil
}
