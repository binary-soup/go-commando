package command

import (
	"github.com/binary-soup/go-command/alert"
	"github.com/binary-soup/go-command/config"
	"github.com/binary-soup/go-command/style"
)

type ConfigCommandBase[T config.Config] struct {
	CommandBase
	config *string
}

func NewConfigCommandBase[T config.Config](name, desc string) ConfigCommandBase[T] {
	cmd := ConfigCommandBase[T]{
		CommandBase: NewCommandBase(name, desc),
	}

	cmd.config = cmd.Flags.String("config", "", "path to a custom config file")
	return cmd
}

func (cmd ConfigCommandBase[T]) LoadConfig() (*T, error) {
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
func (cmd ConfigCommand[T]) Run(args []string) error {
	validate := cmd.Flags.Bool("v", true, "validate the config file")
	cmd.Flags.Parse(args)

	cfg, err := cmd.LoadConfig()
	if err != nil {
		return err
	}

	if *validate {
		err = cmd.validate(*cfg)
	}
	return err
}

func (cmd ConfigCommand[T]) validate(cfg T) error {
	errs, err := cfg.Validate()
	if err != nil {
		return alert.ChainError(err, "error validating config")
	}

	if len(errs) == 0 {
		style.BoldSuccess.Println("Config VALID!")
		return nil
	}

	for i := len(errs) - 1; i >= 0; i-- {
		err = alert.ChainErrorF(err, "%s %s", style.Error.Format("[X]"), errs[i].Error())
	}
	return alert.ChainError(err, "invalid config")
}
