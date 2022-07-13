package main

import (
	"fmt"
	"fpkg/funcs"
	"log"
	"os"
	"text/tabwriter"
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

  w := new(tabwriter.Writer)
  w.Init(os.Stdout, 8, 8, 0, '\t', 0)

  defer w.Flush()

  fmt.Fprintf(w, "\n %s\t%s\t%s\t", "Repository", "Stars", "IsFork")
  fmt.Fprintf(w, "\n %s\t%s\t%s\t", "----------", "-----", "------")
  
  if len(repos.Repositories) == 0 {
    log.Fatal("Package not found.")
  }
  if len(repos.Repositories) < SizeOfList {
    SizeOfList = len(repos.Repositories)
  }

  for i := 0; i < SizeOfList;i++ {
    fmt.Fprintf(w, "\n [%d] %s\t%d\t%s\t", i,
    repos.Repositories[i].GetFullName(),
    repos.Repositories[i].GetStargazersCount(),
    funcs.IsFork(repos.Repositories[i].GetFork()))
  }

  Cli()
}
