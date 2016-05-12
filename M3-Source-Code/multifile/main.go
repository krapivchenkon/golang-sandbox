// multifile project main.go
package main

import (
	// cs "CarStuff"
	"fmt"
)

func main() {
	// c := cs.Car{4, 6}
	// fmt.Println(c)
	// fmt.Println(c.GetDoors())

	t := Truck{2, "full", oneTon}
	fmt.Println(t)
	fmt.Println(t.GetDoors())
}
