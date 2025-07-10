package command

import (
	"github.com/binary-soup/go-command/config"
	"github.com/binary-soup/go-command/style"
)

type ConfigCommandBase[T config.Config] struct {
	CommandBase
	config *string
}

func NewConfigCommandBase[T config.Config](name, desc string) ConfigCommandBase[T] {
	base := NewCommandBase(name, desc)

	return ConfigCommandBase[T]{
		CommandBase: base,
		config:      base.Flags.String("cfg", "", "path to a custom config file"),
	}
}

func (cmd ConfigCommandBase[T]) LoadConfig() (T, error) {
	if *cmd.config != "" {
		return config.LoadCustom[T](*cmd.config)
	}
	return config.LoadDefault[T]()
}

//#########################################

// Command for performing operations with the config file.
type ConfigCommand[T config.Config] struct {
	ConfigCommandBase[T]
}

// Creates a new config command.
func NewConfigCommand[T config.Config]() ConfigCommand[T] {
	return ConfigCommand[T]{
		ConfigCommandBase: NewConfigCommandBase[T]("config", "perform config related tasks"),
	}
}

// Runs the config commands. Run with -h for details.
func (cmd ConfigCommand[T]) Run() error {
	_, err := cmd.LoadConfig()
	if err != nil {
		return err
	}

	style.BoldSuccess.Println("Config VALID!")
	return nil
}
