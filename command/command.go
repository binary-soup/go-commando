// Provides several types for managing multiple commands within a single application.
package command

import (
	"flag"
	"fmt"
	"local/style"
)

// The interface definition of a command.
// Any type that implements the interface can be used by the runner.
type Command interface {
	// Returns the command's name; essentially it's identifier
	GetName() string

	// Prints usage information
	PrintUsage()

	// Runs the command using the provided args and returns any error
	Run(args []string) error
}

// A base struct definition for a command.
// Provides default implementations for some command interface methods.
type CommandBase struct {
	// The name for command. Used by the GetName function
	Name string

	// The description of the command's usage. Used by the PrintUsage function
	Description string

	// The flag set for the command. Please use NewCommandBase to initialize correctly
	Flags *flag.FlagSet
}

// Creates a new command base; mostly used when creating child commands.
// Additionally sets a custom usage function for flag set.
func NewCommandBase(name string, desc string) CommandBase {
	cmd := CommandBase{
		Name:        name,
		Description: desc,
		Flags:       flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.Flags.Usage = func() {
		cmd.PrintUsage()
		style.New(style.Magenta).Println("Options:")
		cmd.Flags.PrintDefaults()
	}

	return cmd
}

// Command interface implementation. Returns the Name field.
func (cmd CommandBase) GetName() string {
	return cmd.Name
}

// Command interface implementation. Prints usage using the command's Name and Description fields.
func (cmd CommandBase) PrintUsage() {
	fmt.Printf("%s | %s\n", style.New(style.Bold, style.Cyan).Format(cmd.Name), cmd.Description)
}
