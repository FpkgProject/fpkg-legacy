package main

import (
	"flag"
	"fmt"
)

func main() {
	installFlag := flag.String("install", "", "install a new package")
	newrepo := flag.String("addrepo", "", "adding a new repo in repoList")
	flag.Parse()
	if *installFlag != "" {
		fmt.Println(*installFlag)
	} else if *newrepo != "" {
		fmt.Println(*newrepo)
	}
}
