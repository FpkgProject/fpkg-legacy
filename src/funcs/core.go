// Core functions / fpkg functions.

package funcs

import (
	"errors"
	"fmt"
	"os"
)

var repoLocale string = "./repoList.txt"

func AddRepo(RepoName string) {
	if _, err := os.Stat(repoLocale); errors.Is(err, os.ErrNotExist) {
		CreateRepoList()
		fmt.Printf("adding the repo %s\n", RepoName)
	} else {
		fmt.Printf("adding the repo %s\n", RepoName)
	}
}

func DelRepo(RepoName string) {
	fmt.Printf("Removing the repo %s\n", RepoName)
}

func InstallPackage(PackageName string) {
	CreateTmpDirectory()
	fmt.Printf("Installing %s\n", PackageName)
	//defer removeTmpDirectory()
}

func UninstallPackage(PackageName string) {
	fmt.Printf("Removing %s\n", PackageName)
}
