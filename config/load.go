package config

import (
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/alert"
	"github.com/binary-soup/go-command/util"
)

var defaultPath = "config.json"

// Set the default config path used by LoadDefault.
//
// Note: the path is relative to the binary, not the execution path.
func SetDefault(path string) {
	defaultPath = path
}

// Load the config from the default path set by SetDefault, or "config.json" if not set.
func LoadDefault[T Config]() (*T, error) {
	path, err := os.Executable()
	if err != nil {
		return nil, alert.ChainError(err, "error finding executable path")
	}
	return LoadCustom[T](filepath.Join(filepath.Dir(path), defaultPath))
}

// Load the config from a custom path.
func LoadCustom[T Config](path string) (*T, error) {
	cfg, err := util.LoadJSON[T]("config", path)
	if err != nil {
		return nil, err
	}

	err = (*cfg).Load()
	if err != nil {
		return nil, alert.ChainError(err, "error loading config")
	}

	return cfg, nil
}
