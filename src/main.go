package main

var (
	MAKEFILENOTFOUND  string = "MakeFile not found in this repository"
	CONFIGURENOTFOUND string = "./configure not found in this repository"
	MAKEFILEexists    bool   = false
	CONFIGUREexists   bool   = false
)

// -------------------------------------------------- Main Function ----------------------------------------
func main() {
	Cli()
}
