package main

import (
	"encoding/json"
	"fmt"
)

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

	text := `[{"A":1,"B":2},{"A":3,"B":4}]`
	var result []*Test
	err := json.Unmarshal([]byte(text), result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}


