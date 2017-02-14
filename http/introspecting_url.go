package main

import (
	"fmt"
	"net/url"
	"strings"
)

// returns last segment of URL
func main() {
	r, e := url.Parse("https://localhost:8080/test/departments")
	if e != nil {
		panic(e)
	}
	p := strings.Split(r.RequestURI(), "/")
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] != "" {
			fmt.Println(p[i])
			break

		}
	}
}
