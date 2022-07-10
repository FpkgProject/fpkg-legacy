// Core functions / fpkg functions.

package funcs

import (
	"fmt"
	"os"
	"errors"
)

func AddRepo(RepoName string) {
	if _, err := os.Stat("./repoList.txt"); errors.Is(err, os.ErrNotExist) {
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

