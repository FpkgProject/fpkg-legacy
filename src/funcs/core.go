package funcs

import (
	"fmt"
)

var (
	tempDirectory string = "tmp"
)

func UninstallPackage(PackageName string) {
	fmt.Printf("Removing %s package from you system", PackageName)
}
