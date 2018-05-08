package main

import (
	"fmt"
	"log"
	"os"
	"paccat/instances"
)

var rawArgs []string

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var subCmd string

	if len(os.Args) > 1 {
		subCmd = os.Args[1]
		if len(os.Args) > 2 {
			rawArgs = os.Args[2:]
		} else {
			rawArgs = nil
		}
	} else {
		subCmd = ""
		rawArgs = nil
	}

	router(subCmd)
}

func router(subCmd string) {
	switch subCmd {
	case "version":
		fmt.Printf("Version: %s\n", Version())
		return
	case "help":
		docHelp()
		return
	case "init":
		instances.Init(rawArgs)
		return
	case "add":
		return
	case "update":
		return
	case "freeze":
		return
	case "destroy":
		return
	case "test":
		return
	case "check":
		return
	case "checkall":
		return
	case "config":
		return
	default:
		echo()
		return
	}
}
