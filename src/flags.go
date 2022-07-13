package main

import (
	"flag"
	"fmt"
	"fpkg/funcs"
)

func init() {
	installFlag := flag.String("install", "", "Install a new package")
	uninstallFlag := flag.String("uninstall", "", "Remove a package from the system")
	flag.Parse()

	if *installFlag != "" {
		searchPackage(*installFlag)
		fmt.Println()
	}
	if *uninstallFlag != "" {
		funcs.UninstallPackage(*uninstallFlag)
	}
}
