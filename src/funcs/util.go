package funcs

import "os"

func CreateRepoList() {
	os.Create("repolist.txt")
}

func CreateTmpDirectory() {
	os.Mkdir("tmp", 0755)
}

func RemoveTmpDirectory() {
	os.Remove("tmp")
}


