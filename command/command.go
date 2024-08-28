package command

import "fmt"

type Command interface {
	GetName() string
	PrintUsage()
	Run() error
}

type CommandBase struct {
	Name        string
	Description string
}

func NewCommandBase(name string, desc string) CommandBase {
	return CommandBase{
		Name:        name,
		Description: desc,
	}
}

func (cmd CommandBase) GetName() string {
	return cmd.Name
}

func (cmd CommandBase) PrintUsage() {
	fmt.Printf("%s | %s\n", cmd.Name, cmd.Description)
}
