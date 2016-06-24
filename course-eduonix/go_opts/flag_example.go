package main

// https://gobyexample.com/command-line-flags

import (
	"flag"
	"fmt"
)

func main() {
	var strPtr string
	flag.StringVar(&strPtr, "s", "def", "usage")

	flag.Parse()

	fmt.Println("flag:", strPtr)
	// go run script.go -s=test

	switch strPtr {
	case "s":
		for n := range sq(gen(5)) {
			fmt.Printf("new el:%v\n", n)
		}
	default:
		fmt.Println("Default case")
	}
}
