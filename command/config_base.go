package command

import (
	"github.com/binary-soup/go-commando/config"
)

type ConfigCommandBase[T config.Config] struct {
	CommandBase
	Config  *string
	Profile *string
}

func NewConfigCommandBase[T config.Config](name, desc string) ConfigCommandBase[T] {
	cmd := ConfigCommandBase[T]{
		CommandBase: NewCommandBase(name, desc),
	}

	cmd.Config = cmd.Flags.String("cfg", "", "path to a custom config file")
	cmd.Profile = cmd.Flags.String("prof", "main", "the config profile")
	return cmd
}

func (cmd ConfigCommandBase[T]) UsingConfig() bool {
	return *cmd.Config != ""
}

func (cmd ConfigCommandBase[T]) GetConfigPath(dataDir string) string {
	if cmd.UsingConfig() {
		return *cmd.Config
	}
	return config.GetProfilePath(dataDir, *cmd.Profile)
}

func (cmd ConfigCommandBase[T]) LoadConfig(dataDir string) (*T, error) {
	return config.Load[T](cmd.GetConfigPath(dataDir))
}
