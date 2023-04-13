package build

import "fmt"

var (
	Version string
	GitHash string
	Build   string
)

func Print() {
	fmt.Printf("Version: %s \n", Version)
	fmt.Printf("Git Hash: %s \n", GitHash)
	fmt.Printf("Build Build: %s \n", Build)
}
