package main

import "fmt"

func docHelp() {
	fmt.Println(`[ PacCat: Arch Linux Developer Package Managaer ฅ'ω'ฅ ]

init [repo_dir]       - Initialize a (git) repository .
help                  - Help menu

add     [package_dir] - Add package index, the package folder should have the same name as the package.
update  [package_dir] - Manually trigger package updates.
freeze  [package_dir] - Freeze package update and remove it from task.
destory [package_dir] - Remove package for local and remote.
test    [package_dir] - Try local packing.
		--quiet, -q   - Quiet Mode.

check   [package_dir] - Check source updates.
checkall              - Get all available updates.

config  [*.conf]      - Setting up the configuration file.
	`)
}
