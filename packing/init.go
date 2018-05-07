package packing

import (
	"fmt"
	"os/exec"
	"strings"
	"bytes"
	"paccat/logger"
)

type DependPackage struct {
	PackageName string
	ContainInfo string
}

var DependList = [2]DependPackage {
	{
		PackageName: "devtools",
		ContainInfo: "local/devtools",
	},
	{
		PackageName: "git",
		ContainInfo: "local/git",
	},
} // todo: get DependList from config file.


func CheckDepends() {
	var containBuf bytes.Buffer
	fmt.Printf("> Checking depends ...\n")

	for item := range DependList {
		cmd := exec.Command("pacman", "-Qs", DependList[item].PackageName)
		buf, err := cmd.Output()

		containBuf.WriteString("Exec: pacman -Qs ")
		containBuf.WriteString(DependList[item].PackageName)
		logger.CheckErr(containBuf.String(), err)

		if strings.Contains(string(buf), DependList[item].ContainInfo) {
			fmt.Printf("===> [pacman] '%s' has been installed.\n", DependList[item].PackageName)
		} else {
			cmd = exec.Command("sudo pacman", "-S", DependList[item].PackageName)
			buf, err = cmd.Output()
			if err != nil {
				logger.CheckErr("Cannot install depends", err)
			}
			fmt.Printf("%s", buf)
		}
	}
	fmt.Printf("> Checking depends pass!\n\n")
}
