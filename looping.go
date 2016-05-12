// Looping project main.go
package main

import (
	"fmt"
)

func main() {
	var carTypes [3]string
	carTypes[0] = "Toyota"
	carTypes[1] = "Ford"
	carTypes[2] = "Nissan"

	i := 0
	for i < len(carTypes) {
		fmt.Println(carTypes[i])
		i++
	}

	fmt.Println("\nfor loop verbose")
	for j := 0; j < len(carTypes); j++ {
		fmt.Println(carTypes[j])
	}

	fmt.Println("\nrange loop")
	for k, value := range carTypes {
		fmt.Println("k=", k, " &value = ", value)
	}

	fmt.Println("\nrange loop ignore key or value")
	for _, value := range carTypes {
		fmt.Println("k=", " &value = ", value)
	}

	//looping through maps
	carDealerInventory := make(map[string]int)
	carDealerInventory["A1 Auto"] = 76
	carDealerInventory["Southwest Auto Mall"] = 112

	carDealerInventory2 := make(map[int]string)
	carDealerInventory2[76] = "A1 Auto"
	carDealerInventory2[112] = "Southwest Auto Mall"

	fmt.Println("\nfirst dictionary loop")
	//this doesn't work for maps
	for m := 0; m < len(carDealerInventory2); m++ {
		fmt.Println(carDealerInventory2[m])
	}
	fmt.Println(carDealerInventory2[76])

	fmt.Println("\nsecond (correct) dictionary loop")
	for key, _ := range carDealerInventory {
		fmt.Println(key, " inventory = ", carDealerInventory[key])
	}
}
