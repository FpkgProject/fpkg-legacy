package funcs

import(
  "github.com/google/go-github/github"
  "context"
)

var client = github.NewClient(nil)

func SearchRepos(reposearch string) (*github.RepositoriesSearchResult, error) {
  repos, _, err := client.Search.Repositories(context.Background(), reposearch, nil)

  return repos, err
}


func GetForkText(fork bool) string {
  if fork {
    return "[Forked]"
  } else {
    return ""
  }
}


func GetDescriptionText(desc string) string {
  if desc == "" {
    return "...No description..."
  } else {
    return desc
  }
}

