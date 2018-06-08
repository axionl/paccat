package container

import (
	"fmt"
	"os"
)

type Container struct {
	Name   string
	Status string
}

var SystemBuildDir = "/var/lib/archbuild/"

func CheckContainerStatus() {
	var container Container
	container.Status = "install" // todo: get container status from systemd and system build path
	container.Name = "paccat_root"

	fmt.Println("> Checking container ...")

	if _, err := os.Stat(SystemBuildDir + container.Name); err == nil {
		fmt.Println("> Container has been installed: ", SystemBuildDir+container.Name)
		container.Status = "update"
	}

	switch container.Status {
	case "install":
		installContainer(container)
		break
	case "update":
		updateContainer()
		break
	}
}
