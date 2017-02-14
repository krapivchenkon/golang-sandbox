package main

import (
	"encoding/json"
	"fmt"
)

type top map[string]interface{}

type s struct {
	Name string `json:"name"`
}

func main() {
	t := make(top)
	t["test"] = "tert"
	t["field"] = &s{Name: "TestNames"}

	v, e := json.Marshal(t)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(v[:]))
}
