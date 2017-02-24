package main

import "fmt"
import "reflect"

type Cat struct {
	name  string
	skill string
}

func main() {

	cat1 := &Cat{name: "Toonces", skill: "driving"}

	// cat1 is dereferenced thus cat2 will be of Cat type and not *Cat as initial cat1 variable
	cat2 := *cat1

	// here cat1 is not dereferenced. Thus, cat3 will point to the same structure
	cat3 := cat1

	// Go is comparing struct fields. structs not modified yet
	fmt.Println(*cat1 == cat2, cat1 == cat3)
	//true true

	// this assingment will change name only of the copy
	cat2.name = "Felix"
	// this assingment will change initial structure
	cat3.name = "Tom"

	fmt.Println(cat1.name, "is good at", cat1.skill, " address: ", cat1)
	fmt.Println(cat2.name, "is good at", cat2.skill, " address: ", &cat2)
	fmt.Println(cat3.name, "is good at", cat3.skill, " address: ", cat3)

	// Go is comparing struct fields. Structs modified
	fmt.Println(*cat1 == cat2, cat1 == cat3)
	//false true

	v1 := reflect.ValueOf(cat1)
	v2 := reflect.ValueOf(&cat2)
	v3 := reflect.ValueOf(cat3)
	fmt.Println(v1.Pointer(), v2.Pointer(), v3.Pointer())
}
