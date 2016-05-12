// Arrays_Slices project main.go
package main

import (
	"fmt"
)

func main() {
	var carTypes [3]string
	carTypes[0] = "Toyota"
	carTypes[1] = "Ford"
	carTypes[2] = "Nissan"
	fmt.Println(carTypes[1])
	fmt.Println("carTypes len = ", len(carTypes))

	carTypes2 := [3]string{"Toyota", "Ford", "Nissan"}
	fmt.Println(carTypes2[0])

	carTypesSlice := []string{"Toyota", "Ford", "Nissan"}
	fmt.Println(carTypesSlice[2])
	carTypesSlice = append(carTypesSlice, "Telsa")
	fmt.Println("carTypesSlice len = ", len(carTypesSlice))

	carTypesSliceMake := make([]string, 3)
	carTypesSliceMake[0] = "Toyota"
	carTypesSliceMake[1] = "Ford"
	carTypesSliceMake[2] = "Nissan"

	fmt.Println("carTypesSlice2= ", carTypesSlice)
	carTypesSlice2 := carTypesSlice[1:3]
	fmt.Println("carTypesSlice2 len = ", len(carTypesSlice2))
	fmt.Println("slice[2:4] = ", carTypesSlice2)

}
