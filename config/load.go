package config

import (
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/alert"
	"github.com/binary-soup/go-command/data"
)

// Calculate the config path from the environment.
// These files are stored relative to the binary, and should be named in the style of {env}.json.
func GetEnvPath(env string) string {
	path, _ := os.Executable()
	return filepath.Join(filepath.Dir(path), env+".json")
}

// Load the config from the given path.
// Additionally calls config.Load and config.Validate before returning the new struct.
func Load[T Config](path string) (T, error) {
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
