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

func AutoPack(packageInfo Package, quiet_mode bool) {
	var containBuf bytes.Buffer
	fmt.Printf("> Making package ...\n")
	fmt.Printf("===> [extra-x86_64-build] Package Name: `%s`", packageInfo.Name)

	os.Chdir(packageInfo.Path)
	cmd := exec.Command("extra-x86_64-build")
	buf, err := cmd.Output()

	containBuf.WriteString("Exec: extra-x86_64-build")
	containBuf.WriteString(packageInfo.Name)
	logger.CheckErr(containBuf.String(), err)

	if !quiet_mode {
		fmt.Printf("===> [extra-x86_64-build] %s", buf)
	}
	fmt.Printf("Building success! \n")
}
