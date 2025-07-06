package sample

import (
	"os"

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
func (cfg SampleConfig) Validate() ([]error, error) {
	errs := []error{}

	if cfg.Path == "" {
		errs = append(errs, alert.Error("path cannot be empty"))
	} else {
		_, err := os.Stat(cfg.Path)
		if os.IsNotExist(err) {
			errs = append(errs, alert.Error("path does not exist"))
		} else if err != nil {
			return nil, err
		}
	}

	if cfg.Count < COUNT_MIN {
		errs = append(errs, alert.ErrorF("count less than min %d", COUNT_MIN))
	} else if cfg.Count > COUNT_MAX {
		errs = append(errs, alert.ErrorF("count more than max %d", COUNT_MAX))
	}

	return errs, nil
}
