package main

import (
	"github.com/spf13/pflag"
)

var InstallFlag string

func init() {
  pflag.StringVarP(&InstallFlag, "install", "i", "", "Install a new package.")
  pflag.StringP("remove", "r", "", "Remove a package.")
  pflag.Parse()

  //UninstallFlag := flag.String("uninstall", "", "Remove a package from the system")

  /*
	if *InstallFlag != "" {
		funcs.InstallPackage(*InstallFlag)
	}
	if *UninstallFlag != "" {
		funcs.UninstallPackage(*UninstallFlag)
	}
  */
}
