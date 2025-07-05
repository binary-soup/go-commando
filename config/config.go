package config

// Interface definition of a Config object.
type Config interface {
	// Runs when the config is first loaded. Useful for initialization related tasks.
	Load() error
	// Validates the config for correctness. Returns a slice of validations errors, and a separate error for any other errors.
	Validate() ([]error, error)
}
