package build

import "fmt"

var (
	Version string
	GitHash string
	Build   string
)

func Print() {
	fmt.Printf("Version: %s \n", Version)
	fmt.Printf("Build Time: %s \n", Build)
	fmt.Printf("Git Hash: %s \n", GitHash)
}
