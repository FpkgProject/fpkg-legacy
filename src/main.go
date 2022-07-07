package main

import (
	"fmt"
)

var PACKAGENOTFOUND = "The package not found in list of repo"
var PACKAGEFOUND = "The package are found in list of repo"

func main() {

	var locked = true
	if locked == false {
		fmt.Printf("%s\n", PACKAGENOTFOUND)
	} else {
		fmt.Printf("%s\n", PACKAGEFOUND)
	}
}
