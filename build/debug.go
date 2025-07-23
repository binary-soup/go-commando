//go:build !release

package build

// Returns whether the project was built using the release tag ("-tags release").
func IsReleaseMode() bool {
	return false
}

// Return the build mode as a string (debug or release).
func GetMode() string {
	return "debug"
}
