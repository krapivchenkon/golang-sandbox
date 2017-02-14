// Maps project main.go
package main

import (
	"fmt"
)

func main() {
	carDealerInventory := make(map[string]int)
	carDealerInventory["A1 Auto"] = 76
	fmt.Println("A1 Auto Inv = ", carDealerInventory["A1 Auto"])
	fmt.Println("len of carDealerInventory = ", len(carDealerInventory))

	carDealerInventory["Southest Auto Mall"] = 112
	fmt.Println("len of carDealerInventory = ", len(carDealerInventory))
	delete(carDealerInventory, "A1 Auto")
	fmt.Println("len of carDealerInventory after delete ", len(carDealerInventory))
	fmt.Println("A1 Auto Inv = ", carDealerInventory["A1 Auto"])
	fmt.Println("Southest Auto Mall Inv = ", carDealerInventory["Southest Auto Mall"])

	carDealerTown := make(map[string]string)
	carDealerTown["A1 Auto"] = "Austin, TX"
	carDealerTown["Southest Auto Mall"] = "Phoenix, AZ"
	fmt.Println("A1 Auto is located in ", carDealerTown["A1 Auto"])

	dict := make(map[string]string)

	val, ok := dict["key1"]

	fmt.Println(val)

	fmt.Println(ok)

	// Check if key exists
	if _, ok := dict["key1"]; ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key is missing")
	}

	// loop over keys
	dict["key2"] = "val2"
	dict["key3"] = "val3"
	dict["key4"] = "val4"

	for key, val := range dict {
		fmt.Println("Key=", key, " Val=", val)
	}
	// remove key from map
	delete(dict, "key1")
}
