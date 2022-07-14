package main

import (
  "strings"
  "github.com/spf13/pflag"
)

var (
  installTmpFlag string
  InstallFlag    []string
  SearchFlagVar  string
)

func init() {
  pflag.StringVarP(&installTmpFlag, "install", "i", "", "Install a new package.")
  pflag.StringP("remove", "r", "", "Remove a package.")
  pflag.StringVarP(&SearchFlagVar, "search", "s", "", "search for a package.")
  pflag.Parse()

  if installTmpFlag != "" {
    InstallFlag = strings.Split(installTmpFlag, "/")
    Install()
  }
}

