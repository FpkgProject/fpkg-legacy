package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	MAKEFILENOTFOUND  string = "The Makefile not is found in this repository"
	CONFIGURENOTFOUND string = "The ./configure not is found in this repository"
	MAKEFILEexists    bool   = false
	CONFIGUREexists   bool   = false
)

// ----------------------------------------------- Util functions -------------------------------------------

func createRepoList() {
	os.Create("repolist.txt")
}

func createTmpDirectory() {
	os.Mkdir("tmp", 0755)
}

func removeTmpDirectory() {
	os.Remove("tmp")
}

// ----------------------------------------------- List of Function in Fpkg ---------------------------------

func addRepo(RepoName string) {
	if _, err := os.Stat("./repoList.txt"); errors.Is(err, os.ErrNotExist) {
		createRepoList()
		fmt.Printf("adding a repo %s\n", RepoName)
	} else {
		fmt.Printf("adding a repo %s\n", RepoName)
	}
}

func delRepo(RepoName string) {
	fmt.Printf("Removing a repo %s\n", RepoName)
}

func installPackage(PackageName string) {
	createTmpDirectory()
	fmt.Printf("Installing a %s\n", PackageName)
	//defer removeTmpDirectory()
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
