// ReadWriteFile project main.go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	//write
	stringContent := string(b) + "\nmore new stuff"
	fmt.Println("write to file", stringContent)
	err = ioutil.WriteFile("output.txt", []byte(stringContent), 0644)
	if err != nil {
		panic(err)
	}
}
