package packing

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"paccat/logger"
)

type Package struct {
	Name    string
	Version string
	Release string
	Path    string
	Url     string
}

func AutoPack(pkg Package, quietMode bool) {
	var containBuf bytes.Buffer

	checkDevTools(quietMode)

	fmt.Println("> Making package ...")
	fmt.Printf("===> [extra-x86_64-build] Package Name: `%s`\n", pkg.Name)

	os.Chdir(pkg.Path)
	cmd := exec.Command("extra-x86_64-build")
	buf, err := cmd.Output()

	containBuf.WriteString("Exec: extra-x86_64-build")
	containBuf.WriteString(pkg.Name)
	logger.CheckErr(containBuf.String(), err)

	if !quietMode {
		fmt.Printf("===> [extra-x86_64-build] %s\n", buf)
	}
	fmt.Printf("> Packaged successfully!\n")
}
