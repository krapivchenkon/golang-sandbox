package main

import "fmt"

func main() {
	defer func() {
		// recover should be called directly in defer func
		// something like:
		// defer func() {
		//        doRecover() //panic is not recovered
		//         }()
		// won't work

		fmt.Println("recovered:", recover())
	}()

	panic("not good")
}
