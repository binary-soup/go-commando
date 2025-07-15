package command

import (
	"github.com/binary-soup/go-command/config"
	"github.com/binary-soup/go-command/style"
)

// A base struct definition for a config command.
// Provides custom flags and helper methods for accepting and loading config files.
type ConfigCommandBase[T config.Config] struct {
	CommandBase
	env    *string
	config *string
}

// Creates a new config command base; mostly used when creating child commands.
// Additionally defines flags for environment and config path.
func NewConfigCommandBase[T config.Config](name, desc string) ConfigCommandBase[T] {
	base := NewCommandBase(name, desc)

	return ConfigCommandBase[T]{
		CommandBase: base,
		env:         base.Flags.String("env", "main", "the config environment"),
		config:      base.Flags.String("cfg", "", "path to a custom config file"),
	}
}

// Load the config file. Uses the custom path if set, else uses the environment.
func (cmd ConfigCommandBase[T]) LoadConfig() (T, error) {
	if *cmd.config != "" {
		return config.LoadCustom[T](*cmd.config)
	}
	return config.LoadEnv[T](*cmd.env)
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
