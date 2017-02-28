package main

import "fmt"

type ctx1 int
type ctx2 int

const (
	key1 ctx1 = 0
	key2 ctx2 = 0
)

// CustomTypesComp Go doesn't perform automatic type casts when comparing values
// even if custom types derived from the same type.
// thus its safe to use your custom type to safely put your values in context package
func CustomTypesComp(a, b interface{}) {
	fmt.Println("Custom types compare:")
	fmt.Println("a == b = ", a == b)
	// Now check if it will be different keys in map.

	d := make(map[interface{}]string)
	d[key1] = "key1"
	d[key2] = "key2"
	fmt.Println("map: ", d)
	fmt.Println("key1=", d[key1], "; key2=", d[key2])
	fmt.Println("=========")
	// Output:
	// Custom types compare:
	// a == b =  false
	// map:  map[0:key1 0:key2]
	// key1= key1 ; key2= key2
}

func main() {
	CustomTypesComp(key1, key2)

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
// type customslice []string
//
// realslice := []string(customslice{"one", "two"})    // convert from custom type to a basic type
// cs customslice = customslice([]string{"one", "two"}) // convert from a basic type to a custom type
//
// //Using custom types2
//
// type Timestamp time.Time
//
// stmp:=Timestamp(time.Now())
//
// rubyDate:=time.Time(stmp).Format(time.RubyDate)
