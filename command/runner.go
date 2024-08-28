package command

import "fmt"

type Runner struct {
	Commands map[string]Command
}

func NewRunner(commands ...Command) Runner {
	cmds := map[string]Command{}
	for _, cmd := range commands {
		cmds[cmd.GetName()] = cmd
	}

	return Runner{
		Commands: cmds,
	}
}

func (r Runner) RunCommand(name string) error {
	cmd, ok := r.Commands[name]
	if !ok {
		return fmt.Errorf("unknown command \"%s\"", name)
	}

	return cmd.Run()
}

func (r Runner) ListCommands() {
	for _, cmd := range r.Commands {
		cmd.PrintUsage()
	}
}
