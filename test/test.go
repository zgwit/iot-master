package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, _ := regexp.Compile(`^(X|D|O)(\d+)$`)
	fmt.Println(re.FindStringSubmatch("D10"))
	//fmt.Println(re.FindAllStringSubmatch("D10", -1))

	bbb()
	bbb()
}

func bbb()  {
	cc := func() {

	}
	fmt.Println(cc, &cc)
}
