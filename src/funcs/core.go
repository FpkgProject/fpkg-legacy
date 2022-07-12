// Core functions / fpkg functions.

package funcs

import (
	"fmt"
)

var tempDirectory string = "tmp"

func InstallPackage(PackageName string) {
	CreateTmpDirectory()
	fmt.Printf("Installing %s\n", PackageName)
	//defer removeTmpDirectory()
}

func UninstallPackage(PackageName string) {
	fmt.Printf("Removing %s\n", PackageName)
}
