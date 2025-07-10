package config

import (
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/alert"
	"github.com/binary-soup/go-command/data"
)

var defaultPath = "config.json"

// Set the default config path used by LoadDefault.
//
// Note: the path is relative to the binary, not the execution path.
func SetDefault(path string) {
	defaultPath = path
}

// Load the config from the default path set by SetDefault, or "config.json" if not set.
func LoadDefault[T Config]() (T, error) {
	var empty T

	path, err := os.Executable()
	if err != nil {
		return empty, alert.ChainError(err, "error finding executable path")
	}
	return LoadCustom[T](filepath.Join(filepath.Dir(path), defaultPath))
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
