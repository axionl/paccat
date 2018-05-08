package local

import (
	"errors"
	"fmt"
	"io/ioutil"
	"paccat/logger"
	"strings"
	"os"
	"os/exec"
)

var gitFlag bool

func Scan(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	logger.CheckErr("Get files", err)
	if len(files) == 0 {
		err := errors.New("please initialized the repo")
		logger.CheckErr("Have no packages here: ", err)
	}

	var packages []string

	for _, item := range files {
		if strings.HasPrefix(item.Name(), ".") {
			if item.Name() == ".git" {
				gitFlag = true
			}
		} else {
			if item.IsDir() {
				packages = append(packages, item.Name())
			}
		}
	}

	fmt.Println("[git] Checking ...")
	initGit(dir)
	
	return packages
}

func initGit(dir string) {
	if !gitFlag {
		os.Chdir(dir)
		cmd := exec.Command("git", "init")

		buf, err := cmd.Output()
		logger.CheckErr("Exec: git initializing:", err)
		fmt.Printf("===> [git] %s", buf)
	} else {
		fmt.Println("===> [git] Git file existed.")
	}
}
