package local

import (
	"os"
	"paccat/logger"
	"bufio"
	"strings"
	"io/ioutil"
	"errors"
	"paccat/packing"
)

func readFile(pkgPath string) (*os.File, *bufio.Scanner) {
	file, err := os.OpenFile(pkgPath, os.O_RDONLY, 0644)
	logger.CheckErr("OpenFile: ", err)
	scanner := bufio.NewScanner(file)

	return file, scanner
}

func check(pkgPath string) bool {
	/* Check PKGBUILD file.*/
	files, err := ioutil.ReadDir(pkgPath)
	logger.CheckErr("Read dir: ", err)
	for _, file := range files{
		if file.Name() == "PKGBUILD" {
			return true
		}
	}
	return false
}

func parse(pkgPath string) (packing.Package, error) {
	var pkg packing.Package

	res := check(pkgPath)
	if res {
		pkg.Path = pkgPath

		os.Chdir(pkgPath)
		file, scanner := readFile("PKGBUILD")
		defer file.Close()

		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "=") {
				str := strings.Split(scanner.Text(), "=")
				switch str[0] {
				case "pkgname":
					pkg.Name = str[1]
				case "pkgver":
					pkg.Version = str[1]
				case "pkgrel":
					pkg.Release = str[1]
				case "url":
					pkg.Url = str[1]
				}
			}
		}
		logger.CheckErr("Scanner:", scanner.Err())
		return pkg, nil
	}
	return pkg, errors.New("Have no PKGBUILD file.")
}