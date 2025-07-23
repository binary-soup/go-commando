//go:build release

package build

func IsReleaseMode() bool {
	return true
}

func GetMode() string {
	return "release"
}
