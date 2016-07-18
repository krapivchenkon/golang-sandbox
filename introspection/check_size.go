package main

import "fmt"
import "unsafe"

func main() {
	var s string
	var i int8
	var c complex128
	s1 := "2342423"
	fmt.Println(unsafe.Sizeof(s))  // prints 8
	fmt.Println(unsafe.Sizeof(c))  // prints 16
	fmt.Println(unsafe.Sizeof(s1)) // prints 16
	fmt.Println(unsafe.Sizeof(i))  // prints 16
}
