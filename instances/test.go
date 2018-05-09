package instances

import (
	"errors"
	"paccat/local"
	"paccat/logger"
	"paccat/packing"
	"strings"
)

func Test(rawArgs []string) {
	var packagePath string
	quietMode := false

	for _, item := range rawArgs {
		if strings.Contains("-q", item) || strings.Contains("--quiet", item) {
			quietMode = true
			if len(rawArgs) > 2 {
				err := errors.New("don't make more than one package at a time")
				logger.CheckErr("Making Package: ", err)
			}
		} else {
			packagePath = item
		}
	}

	if len(packagePath) == 0 {
		packagePath = "."
	}

	pkg, err := local.Parse(packagePath)
	logger.CheckErr("Parse PKGBUILD: ", err)

	packing.AutoPack(pkg, quietMode)
}
