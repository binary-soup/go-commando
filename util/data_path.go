package util

import (
	"os"
	"path/filepath"

	"github.com/binary-soup/go-commando/build"
)

// Get the path for storing user data.
// Of the form <OS config dir>/<data dir>/<build mode>.
//
// Panics if the location cannot be determined.
func UserDataPath(dataDir string) string {
	path, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(path, dataDir, build.GetMode())
}
