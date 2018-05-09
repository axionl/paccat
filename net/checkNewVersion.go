package net

import (
	"fmt"
	"paccat/local"
	"paccat/logger"
	"paccat/packing"
)

func CheckNewVersion(packagePath string) packing.Package {
	pkg, err := local.Parse(packagePath)
	logger.CheckErr("Parse PKGBUILD: ", err)
	fmt.Println(pkg)
	newPackage := nvChecker(pkg)
	return newPackage
}

func nvChecker(pkg packing.Package) packing.Package {
	// todo: get new package version from git, aur or website.
	var newPackage packing.Package

	return newPackage
}
