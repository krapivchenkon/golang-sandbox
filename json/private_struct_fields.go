package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One int
	two string // value will be ignored while decoding
}

// The struct fields starting with lowercase letters
// will not be (json, xml, gob, etc.) encoded, so when
// you decode the structure you'll end up with zero
// values in those unexported fields.

func main() {
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in) //prints main.MyData{One:1, two:"two"}

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) //prints {"One":1}

	var out MyData
	json.Unmarshal(encoded, &out)

	fmt.Printf("%#v\n", out) //prints main.MyData{One:1, two:""}
}
