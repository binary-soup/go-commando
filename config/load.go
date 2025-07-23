package config

import (
	"path/filepath"

	"github.com/binary-soup/go-commando/alert"
	"github.com/binary-soup/go-commando/util"
)

// Calculate the config path from the profile.
// These files are stored in the user data directory, and should be named in the style of {prof}.config.json.
func GetProfilePath(dataDir, prof string) string {
	return filepath.Join(util.UserDataPath(dataDir), prof+".config.json")
}

// Load the config from a custom path.
func Load[T Config](path string) (*T, error) {
	cfg, err := util.LoadJSON[T]("config", path)
	if err != nil {
		return nil, err
	}

	err = (*cfg).Load()
	if err != nil {
		return nil, alert.ChainError(err, "error loading config")
	}

	err = validate(*cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
