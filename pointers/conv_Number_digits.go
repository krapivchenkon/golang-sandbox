package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 123414

	var digits []byte = []byte(strconv.Itoa(i))
	for j, v := range digits {
		fmt.Println(j, "=", string(v))
	}
}
