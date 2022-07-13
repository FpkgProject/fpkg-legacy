package main

import (
  "github.com/IbrahimShahzad/gonfigure"
  "errors"
  "fmt"
  "log"
  "os"
  "strconv"
)

var(
  cfgPath    string = "./config.ini"
  SizeOfList int
)

func newCfg() error {
  newObj := gonfigure.InitialiseINIobj()
  newObj = gonfigure.InsertSection(newObj, "options")
  newObj, err := gonfigure.WriteParameterToSection(newObj, "options", "sizeoflist", "5")
  if err != nil{
    return err
  } /* If we get a error, return to the main function to handle
    * Because the next line will overwrite it.              */
  err = gonfigure.WriteINIFile(newObj, cfgPath)
  return err
}

func init(){
  if _, err := os.Stat(cfgPath); errors.Is(err, os.ErrNotExist){
    fmt.Println("Config file not found, creating one.")
    err := newCfg()
    if err != nil {
      log.Fatal(err)
    }
  }
  cfgfile, err := gonfigure.LoadINI(cfgPath)
  if err != nil {
    log.Fatal("Error loading the config file.\nErr:", err)
  }

  listParam, err := gonfigure.GetParameterValue(cfgfile, "options", "sizeoflist")
  if err != nil {
    fmt.Println("Failed to get SizeOfList value. Using default 5")
    SizeOfList = 5
  } else {
    SizeOfList, err = strconv.Atoi(listParam)
    if err != nil {
      fmt.Println("Failed to get SizeOfList value. Using default 5")
      SizeOfList = 5
    }
  }
}

