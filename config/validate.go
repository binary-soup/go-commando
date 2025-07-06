package config

import (
	"github.com/binary-soup/go-command/alert"
	"github.com/binary-soup/go-command/style"
)

func validate[T Config](cfg T) error {
	errs, err := cfg.Validate()
	if err != nil {
		return alert.ChainError(err, "error validating config")
	}

	if len(errs) == 0 {
		return nil
	}

	for i := len(errs) - 1; i >= 0; i-- {
		err = alert.ChainErrorF(err, "%s %s", style.Error.Format("[X]"), errs[i].Error())
	}
	return alert.ChainError(err, "invalid config")
}
