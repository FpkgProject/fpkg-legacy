package funcs

import "os"

func CreateRepoList() {
	os.Create(*&repoLocale)
}

func CreateTmpDirectory() {
	os.Mkdir(*&tempDirectory, 0755)
}

func RemoveTmpDirectory() {
	os.Remove(*&tempDirectory)
}
