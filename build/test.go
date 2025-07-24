//go:build !prod

package build

// Returns whether the project was built in test mode (ie. not setting "-tags prod").
func IsTest() bool {
	return true
}

// Returns whether the project was built in production mode (ie. setting "-tags prod").
func IsProduction() bool {
	return false
}

// Return the build type as a string ("test" or "production").
func GetType() string {
	return "test"
}
