package main

import (
	"flag"
	"fmt"
)

// ----------------------------------------------- List of Function in Fpkg ---------------------------------

func addRepo(RepoName string) {
	fmt.Printf("adding a repo %s\n", RepoName)
}

func delRepo(RepoName string) {
	fmt.Printf("Removing a repo %s\n", RepoName)
}

func installPackage(PackageName string) {
	fmt.Printf("Installing a %s\n", PackageName)
}

func uninstallPackage(PackageName string) {
	fmt.Printf("uninstalling a %s\n", PackageName)
}

// -------------------------------------------------- CLI ---------------------------------------------------
func cli() {
	installFlag := flag.String("install", "", "install a new package")
	uninstallFlag := flag.String("uninstall", "", "uninstall a package from system")
	addrepo := flag.String("addrepo", "", "adding a new repo in repoList")
	delrepo := flag.String("delrepo", "", "removing a repo in repoList")

	flag.Parse()

	if *installFlag != "" {
		installPackage(*installFlag)
	} else if *addrepo != "" {
		addRepo(*addrepo)
	} else if *uninstallFlag != "" {
		uninstallPackage(*uninstallFlag)
	} else if *delrepo != "" {
		delRepo(*delrepo)
	}
}

// -------------------------------------------------- Main Function ----------------------------------------
func main() {
	cli()
}
