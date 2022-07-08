package main

import (
	"fmt"
)

var PACKAGENOTFOUND = "The package not found in list of repo"
var PACKAGEFOUND = "The package are found in list of repo"
var LOCKED = false

func is_locked(Locked bool) {
	if Locked == true {
		fmt.Printf("%s\n", PACKAGEFOUND)
	} else {
		fmt.Printf("%s\n", PACKAGENOTFOUND)
	}
}

func main() {
	is_locked(LOCKED)
}
