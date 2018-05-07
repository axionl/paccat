package packing

import (
	"fmt"
	"os/exec"
	"strings"
	"bytes"
	"github.com/paccat/error"
)

type DependPackage struct {
	PackageName string
	ContainInfo string
}

type DependList struct {
	pkg []DependPackage
}

func CheckDepends() {
	dependList := &DependList{
		pkg: []DependPackage{
			{
				PackageName: "devtools",
				ContainInfo: "local/devtools",
			},
			{
				PackageName: "git",
				ContainInfo: "local/git",
			},
		},
	} // todo: get dependList from config file.

	var containBuf bytes.Buffer

	for item := range dependList.pkg {
		cmd := exec.Command("pacman", "-Qs", dependList.pkg[item].PackageName)
		buf, err := cmd.Output()

		containBuf.WriteString("Exec: pacman -Qs ")
		containBuf.WriteString(dependList.pkg[item].PackageName)
		error.CheckErr(containBuf.String(), err)

		if strings.Contains(string(buf), dependList.pkg[item].ContainInfo) {
			fmt.Printf("[Pacman] '%s' has been installed.\n", dependList.pkg[item].PackageName)
		} else {
			cmd = exec.Command("sudo pacman", "-S", dependList.pkg[item].PackageName)
			buf, err = cmd.Output()
			if err != nil {
				error.CheckErr("Cannot install depends", err)
			}
			fmt.Printf("%s", buf)
		}
	}
}
