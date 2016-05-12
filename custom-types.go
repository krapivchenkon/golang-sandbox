package main

import (
	"fmt"
)

func main() {

	aCar := car{vehicle{5, 2}, red}
	fmt.Println(aCar)
	// fmt.Println(aCar.getMpg())
	aCar.getMpg()
	fmt.Println(aCar)
}

type vehicle struct {
	mpg           int
	numberOFDoors int
}

type car struct {
	vehicle //anonymous field
	Color
}

func (v *vehicle) getMpg() {
	fmt.Println("This vehicle gets:", v.mpg, " mpg")
}

type Color string

const (
	red   Color = "RED-COLOR"
	black Color = "BLACK-COLOR"
)
