package command

import (
	"github.com/binary-soup/go-commando/config"
)

type ConfigCommandBase[T config.Config] struct {
	CommandBase
	config  *string
	profile *string
}

func NewConfigCommandBase[T config.Config](name, desc string) ConfigCommandBase[T] {
	cmd := ConfigCommandBase[T]{
		CommandBase: NewCommandBase(name, desc),
	}

	cmd.config = cmd.Flags.String("cfg", "", "path to a custom config file")
	cmd.profile = cmd.Flags.String("prof", "main", "the config profile")
	return cmd
}

func (cmd ConfigCommandBase[T]) GetConfigPath(dataDir string) string {
	if *cmd.config != "" {
		return *cmd.config
	}
	return config.GetProfilePath(dataDir, *cmd.profile)
}

func (cmd ConfigCommandBase[T]) LoadConfig(dataDir string) (*T, error) {
	return config.Load[T](cmd.GetConfigPath(dataDir))
}
