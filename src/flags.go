package main

import (
	"flag"
	"fpkg/funcs"
)

func Cli() {
	installFlag := flag.String("install", "", "Install a new package")
	uninstallFlag := flag.String("uninstall", "", "Remove a package from the system")
	addrepo := flag.String("addrepo", "", "Add a new repo to the repoList")
	delrepo := flag.String("delrepo", "", "Remove a repo to the repoList")

	flag.Parse()

	if *installFlag != "" {
		funcs.InstallPackage(*installFlag)
	} else if *addrepo != "" {
		funcs.AddRepo(*addrepo)
	} else if *uninstallFlag != "" {
		funcs.UninstallPackage(*uninstallFlag)
	} else if *delrepo != "" {
		funcs.DelRepo(*delrepo)
	}
}
