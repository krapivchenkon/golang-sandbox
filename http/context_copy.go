package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	r, e := http.NewRequest("GET", "/path", nil)
	check(e)

	c := r.Context()
	c1 := context.WithValue(c, "KEY", "TEST")
	r = r.WithContext(c1)
	fmt.Println(r.Context())

	c2 := context.WithValue(c1, "KEY1", "TEST1")
	r = r.WithContext(c2)

	fmt.Println(r.Context())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
