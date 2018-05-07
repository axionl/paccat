package main

import (
	"fmt"
	"paccat/packing"
	"log"
	"os"
)

var rawArgs []string

func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Welcome to PacCat! ฅ'ω'ฅ")
	fmt.Printf("Version: %s\n\n", Version())
	packing.CheckDepends()

	var subCmd string

	if len(os.Args) >= 2 {
		rawArgs = os.Args[1:]
		subCmd = os.Args[1]
	} else {
		rawArgs = nil
		subCmd = "help"
	}

	router(subCmd)
}

func router(subCmd string) {
	switch subCmd {
	case "version":
		fmt.Printf("ฅ'ω'ฅ Version: %s\n", Version())
		return
	case "help":
		docHelp()
		return
	case "init":
		//init(os.Args[2])
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
	//case "echo":
	//	echo(rawArgs[1:]) // test route
	//	return
	}
}

func Version() string {
	version := "0.0.1" // todo: get version from config file.
	return version
}

func docHelp() {
	fmt.Printf("\n")
	fmt.Println(`=== PacCat: Arch Linux Developer Package Managaer. ===

   init [repo_dir] - Initialize a (git) repository .
   help - Help menu

   add [package_dir] - Add package index, the package folder should have the same name as the package.
   update [package_dir] - Manually trigger package updates.
   freeze [package_dir] - Freeze package update and remove it from task.
   destory [package_dir] - Remove package for local and remote.
   test [package_dir] - Try local packing.

   check [package_dir] - Check source updates.
   checkall - Get all available updates.

   config [*.conf] - Setting up the configuration file.
	`)
}

//
//func echo(rawArgs []string){
//	for item := range rawArgs{
//		fmt.Println(rawArgs[item])
//	}
//}