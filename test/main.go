package main

import "fmt"

func main() {

	var a []string

	for i := 0; i < 10; i++ {
		a = append(a, "abc")
	}

	fmt.Println(a)
}
