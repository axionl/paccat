package instances

import (
	"errors"
	"fmt"
	"paccat/logger"
	"paccat/net"
	"paccat/packing"
)

func Test(rawArgs []string) {
	if rawArgs == nil {
		dir = "."
	} else if len(rawArgs) == 1 {
		dir = rawArgs[0]
	} else {
		err := errors.New("don't make more than one package at a time")
		logger.CheckErr("Making Package: ", err)
	}
	fmt.Println(dir)
	pkg := net.CheckNewVersion(dir)
	packing.AutoPack(pkg, true)
}
