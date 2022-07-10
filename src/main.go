package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	MAKEFILENOTFOUND  string = "MakeFile not found in this repository"
	CONFIGURENOTFOUND string = "./configure not found in this repository"
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

// ----------------------------------------------- List of fpkg functions -----------------------------------

func addRepo(RepoName string) {
	if _, err := os.Stat("./repoList.txt"); errors.Is(err, os.ErrNotExist) {
		createRepoList()
		fmt.Printf("adding the repo %s\n", RepoName)
	} else {
		fmt.Printf("adding the repo %s\n", RepoName)
	}
}

func delRepo(RepoName string) {
	fmt.Printf("Removing the repo %s\n", RepoName)
}

func installPackage(PackageName string) {
	createTmpDirectory()
	fmt.Printf("Installing %s\n", PackageName)
	//defer removeTmpDirectory()
}

func uninstallPackage(PackageName string) {
	fmt.Printf("Removing %s\n", PackageName)
}

// -------------------------------------------------- CLI ---------------------------------------------------
func cli() {
	installFlag := flag.String("install", "", "Install a new package")
	uninstallFlag := flag.String("uninstall", "", "Remove a package from the system")
	addrepo := flag.String("addrepo", "", "Add a new repo to the repoList")
	delrepo := flag.String("delrepo", "", "Remove a repo to the repoList")

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
