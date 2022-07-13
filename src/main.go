package main

import (
  "fmt"
  "fpkg/funcs"
  "log"
)


func main(){
  var (
    reposcan string
  )

  fmt.Printf(":: Search for a package: ")
  fmt.Scanln(&reposcan)

  repos, err := funcs.SearchRepos(reposcan)
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
    fmt.Printf("\n\n[%d] -> %s",
      i, repos.Repositories[i].GetFullName())

    fmt.Printf("\n\t%s",
      funcs.GetDescriptionText(repos.Repositories[i].GetDescription()))

    fmt.Printf("\n\tStars: %d %s",
      repos.Repositories[i].GetStargazersCount(),
      funcs.GetForkText(repos.Repositories[i].GetFork()))
  }

  Cli()
}
