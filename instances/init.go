package instances

import (
	"errors"
	"paccat/logger"
	"paccat/local"
)

var dir string
var gitFlag bool

func Init(rawArgs []string) {
	if rawArgs == nil {
		dir = "."
	} else if len(rawArgs) == 1{
		dir = rawArgs[0]
	} else {
		err := errors.New("do not initialize multiple directories")
		logger.CheckErr("Initialized repository: ", err)
	}
	local.Scan(dir)

	//packing.CheckDevTools()
}