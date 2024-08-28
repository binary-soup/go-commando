package command

import (
	"flag"
	"fmt"
)

type Command interface {
	GetName() string
	PrintUsage()
	Run(args []string) error
}

type CommandBase struct {
	Name        string
	Description string
	Flags       *flag.FlagSet
}

func NewCommandBase(name string, desc string) CommandBase {
	cmd := CommandBase{
		Name:        name,
		Description: desc,
		Flags:       flag.NewFlagSet(name, flag.ExitOnError),
	}

	cmd.Flags.Usage = func() {
		cmd.PrintUsage()
		fmt.Println("Options:")
		cmd.Flags.PrintDefaults()
	}

	return cmd
}

func (cmd CommandBase) GetName() string {
	return cmd.Name
}

func (cmd CommandBase) PrintUsage() {
	fmt.Printf("%s | %s\n", cmd.Name, cmd.Description)
}
