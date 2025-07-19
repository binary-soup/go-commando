package command

import (
	"flag"

	"github.com/binary-soup/go-commando/config"
	"github.com/binary-soup/go-commando/data"
	"github.com/binary-soup/go-commando/prompt"
	"github.com/binary-soup/go-commando/style"
)

// A base struct definition for a config command.
// Provides custom flags and helper methods for accepting and loading config files.
type ConfigCommandBase[T config.Config] struct {
	CommandBase
	flags *configBaseFlags
}

type configBaseFlags struct {
	Env    *string
	Config *string
}

func (f *configBaseFlags) Set(flags *flag.FlagSet) {
	f.Env = flags.String("env", "main", "the config environment")
	f.Config = flags.String("cfg", "", "path to a custom config file")
}

// Creates a new config command base; mostly used when creating child commands.
// Additionally defines flags for environment and config path.
func NewConfigCommandBase[T config.Config](name, desc string, flags ...Flags) ConfigCommandBase[T] {
	f := new(configBaseFlags)

	return ConfigCommandBase[T]{
		CommandBase: NewCommandBase(name, desc, append(flags, f)...),
		flags:       f,
	}
}

// Get the config path. Uses the custom path if set, else uses the environment.
func (cmd ConfigCommandBase[T]) GetConfigPath() string {
	if *cmd.flags.Config != "" {
		return *cmd.flags.Config
	}
	return config.GetEnvPath(*cmd.flags.Env)
}

// Load the config file. Uses the custom path if set, else uses the environment.
func (cmd ConfigCommandBase[T]) LoadConfig() (T, error) {
	return config.Load[T](cmd.GetConfigPath())
}

//#########################################

// Command for performing operations with the config file.
type ConfigCommand[T config.Config] struct {
	ConfigCommandBase[T]
	flags *configFlags
}

type configFlags struct {
	New *bool
}

func (f *configFlags) Set(flags *flag.FlagSet) {
	f.New = flags.Bool("new", false, "create a new config file using the set path (env or cfg)")
}

// Creates a new config command.
func NewConfigCommand[T config.Config]() ConfigCommand[T] {
	flags := new(configFlags)

	return ConfigCommand[T]{
		ConfigCommandBase: NewConfigCommandBase[T]("config", "perform config related tasks", flags),
		flags:             flags,
	}
}

// Runs the config commands. Run with -h for details.
func (cmd ConfigCommand[T]) Run() error {
	if *cmd.flags.New {
		return cmd.newConfig()
	}

	_, err := cmd.LoadConfig()
	if err != nil {
		return err
	}

	style.BoldSuccess.Println("Config VALID!")
	return nil
}

func (cmd ConfigCommand[T]) newConfig() error {
	path := cmd.GetConfigPath()

	if prompt.New().ConfirmOverwrite("Config", path) {
		return data.SaveJSON("config", new(T), path)
	}
	return nil
}
