package main

import (
	"flag"
)

func init() {
	installFlag := flag.String("install", "", "Install a new package")
	//UninstallFlag := flag.String("uninstall", "", "Remove a package from the system")
	flag.Parse()

  /*
	if *InstallFlag != "" {
		funcs.InstallPackage(*InstallFlag)
	}
	if *UninstallFlag != "" {
		funcs.UninstallPackage(*UninstallFlag)
	}
  */
}
