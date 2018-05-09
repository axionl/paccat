package instances

import (
	"errors"
	"paccat/local"
	"paccat/logger"
)

func Init(rawArgs []string) {
	var dir string

	if rawArgs == nil {
		dir = "."
	} else if len(rawArgs) == 1 {
		dir = rawArgs[0]
	} else {
		err := errors.New("do not initialize multiple directories")
		logger.CheckErr("Initialized repository: ", err)
	}
	local.Scan(dir)
}
