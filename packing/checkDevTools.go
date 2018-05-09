package packing

import (
	"bytes"
	"fmt"
	"os/exec"
	"paccat/logger"
	"strings"
)

type devTool struct {
	PackageName string
	ContainInfo string
}

var devTools = []devTool{
	{
		PackageName: "devtools",
		ContainInfo: "local/devtools",
	},
	{
		PackageName: "git",
		ContainInfo: "local/git",
	},
} // todo: get devTools from config file.

func checkDevTools() {
	var containBuf bytes.Buffer
	fmt.Printf("> Checking depends ...\n")

	for _, item := range devTools {
		cmd := exec.Command("pacman", "-Qs", item.PackageName)
		buf, err := cmd.Output()

		containBuf.WriteString("Exec: pacman -Qs ")
		containBuf.WriteString(item.PackageName)
		logger.CheckErr(containBuf.String(), err)

		if strings.Contains(string(buf), item.ContainInfo) {
			fmt.Printf("===> [pacman] '%s' has been installed.\n", item.PackageName)
		} else {
			cmd = exec.Command("sudo pacman", "-S", item.PackageName)
			buf, err = cmd.Output()
			if err != nil {
				logger.CheckErr("Cannot install depends", err)
			}
			fmt.Printf("%s", buf)
		}
	}
	fmt.Printf("> Checking depends pass!\n\n")
}
