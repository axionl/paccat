package container

import (
	"fmt"
	"os"
	"os/exec"
	"paccat/logger"
)

func installContainer(container Container) {
	os.MkdirAll(SystemBuildDir+container.Name+"/basic", 0755)
	fmt.Println("==> [pacscripts] Installing the packages `base` & `base-devel`")
	cmd := exec.Command("pacstrap", SystemBuildDir+container.Name+"/basic", "base", "base-devel")
	buf, err := cmd.Output()
	logger.CheckErr("Exec: `pacstrap` ", err)
	fmt.Printf("%s", buf)
	fmt.Println("> Container install finished!")
}

func updateContainer() {
	fmt.Println("==> [pacscripts] Updating the container...")
}
