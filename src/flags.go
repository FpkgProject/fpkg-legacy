package main

import (
  "fmt"
  "os"
  "github.com/spf13/pflag"
)

var InstallFlag string

func init() {
  pflag.StringVarP(&InstallFlag, "install", "i", "", "Install a new package.")
  pflag.StringP("remove", "r", "", "Remove a package.")
  pflag.Parse()

  if len(os.Args) < 2 {
    fmt.Println("\033[1;31m" + "error:" + Color.Reset, "no operation specified.\noperations:")
    pflag.PrintDefaults()
    os.Exit(1)
  }

  //UninstallFlag := flag.String("uninstall", "", "Remove a package from the system")

  /*
  if *InstallFlag != "" {
    funcs.InstallPackage(*InstallFlag)
  }
  if *UninstallFlag != "" {
    funcs.UninstallPackage(*UninstallFlag)
  }*/
}
