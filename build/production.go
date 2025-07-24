//go:build prod

package build

func IsTest() bool {
	return false
}

func IsProduction() bool {
	return true
}

func GetType() string {
	return "production"
}
