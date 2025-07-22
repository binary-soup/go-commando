package command

import (
	"flag"
	"fmt"

	"github.com/binary-soup/go-commando/style"
)

// Interface for defining flags for CommandBase.
type Flags interface {
	// Set the desired flags using the flag set. Called during SubmitArgs.
	Set(flags *flag.FlagSet)
}

// A base struct definition for a command.
// Provides default implementations for some command interface methods.
type CommandBase struct {
	flags []Flags

	// The name for command. Used by the GetName function.
	Name string

	// The description of the command's usage. Used by the PrintUsage function.
	Description string
}

// Creates a new command base; mostly used when creating child commands.
func NewCommandBase(name string, desc string, flags ...Flags) CommandBase {
	return CommandBase{
		flags:       flags,
		Name:        name,
		Description: desc,
	}
}

// Command interface implementation. Returns the Name field.
func (cmd CommandBase) GetName() string {
	return cmd.Name
}

// Command interface implementation. Prints usage using the command's Name and Description fields.
func (cmd CommandBase) PrintUsage() {
	fmt.Printf("%s | %s\n", style.New(style.Bold, style.Cyan).Format(cmd.Name), cmd.Description)
}

// Create a new flag set and parse the submitted args.
// Calls Flags.Set to set the flags beforehand.
func (cmd CommandBase) SubmitArgs(args []string) error {
	flags := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	flags.Usage = func() {
		cmd.PrintUsage()
		style.New(style.Magenta).Println("Options:")
		flags.PrintDefaults()
	}

	for _, f := range cmd.flags {
		if f != nil {
			f.Set(flags)
		}
	}

	return flags.Parse(args)
}
