package packing

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"paccat/container"
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
	// todo: auto pack without archlinux devel-tools.
	var containBuf bytes.Buffer

	checkDevTools(quietMode)

	requireEUID0()
	container.CheckContainerStatus()

	fmt.Println("> Making package ...")
	fmt.Printf("===> [extra-x86_64-build] Package Name: `%s`\n", pkg.Name)

	os.Chdir(pkg.Path)
	cmd := exec.Command("extra-x86_64-build")
	buf, err := cmd.Output()

	containBuf.WriteString("Exec: extra-x86_64-build ")
	containBuf.WriteString(pkg.Name)
	containBuf.WriteString("\nYou may need to reinstall container by `extra-x86_64-build`")
	logger.CheckErr(containBuf.String(), err)

	if !quietMode {
		fmt.Printf("===> [extra-x86_64-build] %s\n", buf)
	}
	fmt.Printf("> Packaged successfully!\n")
}

func requireEUID0() {
	if os.Geteuid() != 0 {
		log.Fatalln("need to be root")
	}
}
