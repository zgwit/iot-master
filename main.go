package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	A int
	B int
}

type Test2 struct {
	Test []*Test
}

func main() {

	text := `{"Test":[{"A":1,"B":2}]}`
	var result Test2
	err := json.Unmarshal([]byte(text), &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.Test[0])
}


