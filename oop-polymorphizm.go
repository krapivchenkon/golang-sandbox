package main

import "fmt"

type Hobby interface {
	myStereotype() string
}

type Human struct {
}

func (h Human) myStereotype() string {
	return "I'm a Human, only an abstract concept, and I can have no hobby."
}

type Man struct {
	Human //anonymous class to inherit Human behavior
}

func (m Man) myStereotype() string {
	return "I'm a Man and I'm going fishing."
}

type Woman struct {
	Human //anonymous class to inherit Human behavior
}

func (m Woman) myStereotype() string {
	return "I'm a Woman and I'm going shopping."
}

type Dog struct {
	//does not inherit any other type
}

func (m Dog) myStereotype() string {
	return "bow bow bow, I'm chasing sticks."
}

func main() {
	h := new(Human)
	m := new(Man)
	w := new(Woman)
	d := new(Dog)

	//an array of hobby instances - we don’t need to know whether human or dog
	hobbyArr := [...]Hobby{h, m, w, d} //array of 3 Humans and 1 dog.
	for n, _ := range hobbyArr {

		fmt.Println("My hobby?  Well,", hobbyArr[n].myStereotype()) //appears as Hobby type, but behavior changes depending on actual instance

	}
}
