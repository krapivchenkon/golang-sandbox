// Conditionals project main.go
package main

import (
	"fmt"
)

func main() {
	carLotA := 11
	carLotB := 29

	if carLotB <= carLotA {
		fmt.Println("car lot B is greater than A")
	} else {
		fmt.Println("condition failed - fall through")
	}

	if carLotA > 15 {
		fmt.Println("value is greater than 15")
	} else if carLotA > 9 {
		fmt.Println("value is greater than 9")
	} else {
		fmt.Println("default case")
	}

	switch carLotA {
	case 15:
		fmt.Println("case value is 15")
	case 9, 11, 12:
		fmt.Println("case value is 9 or 11 or 12--", carLotA)
	default:
		fmt.Println("default case--")
	}
}
