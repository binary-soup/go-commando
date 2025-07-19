// Provides sample commands used to demonstrate the command package.
package sample

import (
	"flag"
	"fmt"

	"github.com/binary-soup/go-commando/alert"
	"github.com/binary-soup/go-commando/command"
	"github.com/binary-soup/go-commando/style"
)

// Sample hello command for printing "Hello" to the console.
type HelloCommand struct {
	command.CommandBase
	flags *helloFlags
}

type helloFlags struct {
	Name *string
}

func (f *helloFlags) Set(flags *flag.FlagSet) {
	f.Name = flags.String("name", "", "name to use when saying hello")
}

// Creates a new HelloCommand.
func NewHelloCommand() HelloCommand {
	flags := new(helloFlags)

	return HelloCommand{
		CommandBase: command.NewCommandBase("hello", "prints hello {name} to the console", flags),
		flags:       flags,
	}
}

// Run the Hello commands. See usage string for details.
func (cmd HelloCommand) Run() error {
	if *cmd.flags.Name == "" {
		return alert.Error("\"name\" cannot be empty")
	}

	fmt.Printf("Hello %s!\n", style.New(style.Bold, style.Yellow).Format(*cmd.flags.Name))
	return nil
}
