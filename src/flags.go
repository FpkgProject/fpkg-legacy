package main

import (
	"flag"
	"fpkg/funcs"
)

func init() {
	installFlag := flag.String("install", "", "Install a new package")
	uninstallFlag := flag.String("uninstall", "", "Remove a package from the system")
	flag.Parse()

	if *installFlag != "" {
		searchPackage(*installFlag)
	}
	if *uninstallFlag != "" {
		funcs.UninstallPackage(*uninstallFlag)
	}
}
