package config

import (
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/alert"
	"github.com/binary-soup/go-command/data"
)

// Load the config by environment.
// These files are stored relative to the binary, and should be named in the style of {name}.json.
func LoadEnv[T Config](name string) (T, error) {
	var empty T

	path, err := os.Executable()
	if err != nil {
		return empty, alert.ChainError(err, "error finding executable path")
	}
	return LoadCustom[T](filepath.Join(filepath.Dir(path), name+".json"))
}

// Load the config from a custom path.
func LoadCustom[T Config](path string) (T, error) {
	cfg, err := data.LoadJSON[T]("config", path)
	if err != nil {
		return cfg, err
	}

	err = cfg.Load()
	if err != nil {
		return cfg, alert.ChainError(err, "error loading config")
	}

	err = validate(cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
