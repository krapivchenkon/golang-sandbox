package main

import (
	"fmt"
)

func main() {
	aCar := car{vehicle{5, 2}}
	fmt.Println(aCar)
	// fmt.Println(aCar.getMpg())
	aCar.getMpg()
	fmt.Println(*aCar)
}

type vehicle struct {
	mpg           int
	numberOFDoors int
}

type car struct {
	vehicle //anonymous field
}

func (v *vehicle) getMpg() {
	fmt.Println("This vehicle gets:", v.mpg, " mpg")
}

// he first important thing to note is that since we have defined Kitchen
// to be an anonymous field, it allows us to access its members as
// if they were members of the encompassing class.

// The second important thing to note is that the composed field is still
// available to be accessed, but by its type name. So in this case the anonymous
// field for Kitchen has to be accessed as h.Kitchen. If you want to print
// the number of plates in the kitchen, then do this: fmt.Println(h.Kitchen.numOfPlates).
