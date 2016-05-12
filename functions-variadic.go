//Variadic Functions
package main

import (
	"fmt"
)

func main() {
	d1 := "A1 Auto"
	d2 := "Discount Auto"
	d3 := "Riverside Automart"

	printDealers(d1, d2, d3)
	fmt.Println("\n")
	dealers := []string{d1, d2, d3}
	printDealers(dealers...)
}

func printDealers(dealers ...string) {
	for _, dealerName := range dealers {
		fmt.Println(dealerName)
	}
}
