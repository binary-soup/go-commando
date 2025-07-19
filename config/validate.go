package config

import (
	"github.com/binary-soup/go-commando/alert"
	"github.com/binary-soup/go-commando/style"
)

func validate[T Config](cfg T) error {
	verrs := cfg.Validate()

	if len(verrs) == 0 {
		return nil
	}
	var err error

	for i := len(verrs) - 1; i >= 0; i-- {
		err = alert.ChainErrorF(err, "%s %s", style.Error.Format("[X]"), verrs[i].Error())
	}
	return alert.ChainError(err, "invalid config")
}
