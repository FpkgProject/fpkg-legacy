package funcs

import "os"

func CreateTmpDirectory() {
	os.Mkdir(*&tempDirectory, 0755)
}

func RemoveTmpDirectory() {
	os.Remove(*&tempDirectory)
}
