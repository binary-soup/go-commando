package command

import "fmt"

// Runner stores multiple commands and provides methods for running by name.
type Runner struct {
	// The map of commands. Please use NewRunner to initialize correctly
	Commands map[string]Command
}

// Creates a new runner using the provided commands.
func NewRunner(commands ...Command) Runner {
	cmds := map[string]Command{}
	for _, cmd := range commands {
		cmds[cmd.GetName()] = cmd
	}

	return Runner{
		Commands: cmds,
	}
}

// Run the command that matches the provided name (case sensitive) with the provided arguments.
// Returns any run errors, or an error if the command was not found.
func (r Runner) RunCommand(name string, args []string) error {
	cmd, ok := r.Commands[name]
	if !ok {
		return fmt.Errorf("unknown command \"%s\"", name)
	}

	return cmd.Run(args)
}

// Print the usage for all commands to the console.
func (r Runner) ListCommands() {
	for _, cmd := range r.Commands {
		cmd.PrintUsage()
	}
}
