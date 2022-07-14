package main

import (
  "fmt"
  "fpkg/funcs"
  "log"
  "os"
  "github.com/go-git/go-git/v5"
)

func Install(){
  url, err := funcs.GetRepoURL(InstallFlag[0], InstallFlag[1])
  if err != nil {
    fmt.Println(Color.Red + "Error:" + Color.Reset, "package not found.")
  }

  git.PlainClone("./root/var/cache/" + InstallFlag[0] + "_" + InstallFlag[1], false, &git.CloneOptions{
    URL:      url,
    Progress: os.Stdout,
  })
}

func SearchFunction(){
repos, err := funcs.SearchRepos(SearchFlagVar)
  if err != nil {
    log.Fatal("Cannot get the repositories list, please check network connection.")
  }

  if len(repos.Repositories) == 0 {
    log.Fatal("Package not found.")
  }

  if len(repos.Repositories) < SizeOfList {
    SizeOfList = len(repos.Repositories)
  }

  for i := 0; i < SizeOfList;i++ {
    fmt.Printf("\n%s%s\033[m",
    Color.Cyan, repos.Repositories[i].GetFullName())

    fmt.Printf("\n\t%s",
    funcs.GetDescriptionText(repos.Repositories[i].GetDescription()))

    fmt.Printf("\n\t%sStars: %d %s\033[m",
    Color.Yellow,
    repos.Repositories[i].GetStargazersCount(),
    funcs.GetForkText(repos.Repositories[i].GetFork()))
  }
  fmt.Println("\n :: Found", repos.GetTotal(), "correspondent packages.")
}

func main(){
  if len(os.Args) < 2 {
    fmt.Println(Color.Bold + Color.Red + "error:" + Color.Reset, "no operation specified. (use -h for help)")
  }
  if SearchFlagVar != "" {
    SearchFunction()
  }
}

