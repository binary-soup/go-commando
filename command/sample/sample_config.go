package sample

import (
	"github.com/binary-soup/go-command/alert"
)

const (
	COUNT_MIN = 5
	COUNT_MAX = 100
)

// A sample config object.
type SampleConfig struct {
	Path  string `json:"path"`
	Count int    `json:"count"`
}

// Implements Config interface Load method.
func (SampleConfig) Load() error {
	return nil
}

// Implements Config interface Validate method.
// Provides example implementation for writing validations.
func (cfg SampleConfig) Validate() []error {
	errs := []error{}

	if cfg.Path == "" {
		errs = append(errs, alert.Error("path cannot be empty"))
	}

	if cfg.Count < COUNT_MIN {
		errs = append(errs, alert.ErrorF("count less than min %d", COUNT_MIN))
	} else if cfg.Count > COUNT_MAX {
		errs = append(errs, alert.ErrorF("count more than max %d", COUNT_MAX))
	}

	return errs
}
