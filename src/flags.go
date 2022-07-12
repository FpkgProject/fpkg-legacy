package main

import (
	"flag"
	"fpkg/funcs"
)

func Cli() {
	installFlag := flag.String("install", "", "Install a new package")
	uninstallFlag := flag.String("uninstall", "", "Remove a package from the system")
	flag.Parse()

	if *installFlag != "" {
		funcs.InstallPackage(*installFlag)
	}
	if *uninstallFlag != "" {
		funcs.UninstallPackage(*uninstallFlag)
	}
}
