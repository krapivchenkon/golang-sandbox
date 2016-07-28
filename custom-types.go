package main

import (
	"fmt"
	"time"
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


// Using custom types
type customslice []string

realslice := []string(customslice{"one", "two"})    // convert from custom type to a basic type
cs customslice = customslice([]string{"one", "two"}) // convert from a basic type to a custom type

//Using custom types2

type Timestamp time.Time

stmp:=Timestamp(time.Now())

rubyDate:=time.Time(stmp).Format(time.RubyDate)