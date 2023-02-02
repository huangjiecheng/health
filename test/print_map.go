package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string
	Age  int64
}

func main() {
	a := make(map[string]A)
	a["a"] = A{
		Name: "haha",
		Age:  10,
	}
	jsonStr, _ := json.Marshal(a)
	fmt.Println(string(jsonStr))
}
