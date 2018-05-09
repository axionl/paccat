package instances

import (
	"errors"
	"paccat/logger"
	"paccat/net"
	"paccat/packing"
)

func Update(rawArgs []string) {
	var packagePath string

	if rawArgs == nil {
		packagePath = "."
	} else if len(rawArgs) == 1 {
		packagePath = rawArgs[0]
	} else {
		err := errors.New("don't make more than one package at a time")
		logger.CheckErr("Making Package: ", err)
	}
	pkg := net.CheckNewVersion(packagePath)
	packing.AutoPack(pkg, true)
}
