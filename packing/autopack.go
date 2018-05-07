package packing

import (
	"os/exec"
	"bytes"
	"fmt"
	"paccat/error"
	"os"
)

type Package struct {
	PackageName string
	PackageVersion string
	PackagePath string
}

func AutoPack(packageInfo Package, quiet_mode bool) {
	var containBuf bytes.Buffer
	fmt.Printf("> Making package ...\n")
	fmt.Printf("===> [extra-x86_64-build] Package Name: `%s`", packageInfo.PackageName)

	os.Chdir(packageInfo.PackagePath)
	cmd := exec.Command("extra-x86_64-build")
	buf, err := cmd.Output()

	containBuf.WriteString("Exec: extra-x86_64-build")
	containBuf.WriteString(packageInfo.PackageName)
	error.CheckErr(containBuf.String(), err)

	if !quiet_mode {
		fmt.Printf("===> [extra-x86_64-build] %s", buf)
	}
	fmt.Printf("Building success! \n")
}
