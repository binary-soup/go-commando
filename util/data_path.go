package util

import (
	"os"
	"path/filepath"

	"github.com/binary-soup/go-commando/build"
)

// Get the path for storing user data.
// Panics if the location cannot be determined.
//
// For Test builds, this is {OS config dir}/{data dir}_test.
// For Production builds, this is just {OS config dir}/{data dir}.
func UserDataPath(dataDir string) string {
	path, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(path, dataDir)

	if build.IsProduction() {
		return path
	} else {
		return path + "_test"
	}
}
