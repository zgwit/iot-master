package main

import "fmt"

type Test struct {
	A int
	B int
}

func main() {
	aa := []Test{{1,2},{3,4}}
	for _, v := range aa {
		v.A++
		v.B--
	}
	fmt.Println(aa)
	//[{1 2} {3 4}] 坑爹的复制
}


