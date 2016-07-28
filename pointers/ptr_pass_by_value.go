package main

import ()

type User struct {
	Name string
}

// sample demontrates that even pointer pass by value copy

func main() {
	u := &User{Name: "Leto"}
	println(u.Name)
	Modify(u)
	println(u.Name)
}

func Modify(u *User) {
	// u.Name = "Paul"
	u = &User{Name: "Paul"}
}

// prints
// Leto
// Leto

// when argument passed by value itself
// func main() {
//   u := User{Name: "Leto"}
//   println(u.Name)
//   Modify(u)
//   println(u.Name)
// }

// func Modify(u User) {
//   u.Name = "Duncan"
// }

// However we can dereference pointer to find its original
func main() {
	u := &User{Name: "Leto"}
	fmt.Println(u.Name)
	Modify(&u)
	fmt.Println(u.Name)
}

func Modify(u **User) {
	*u = &User{Name: "Bob"}
}

// or like this
func main() {
	u := &User{Name: "Leto"}
	fmt.Println(u.Name)
	Modify(u) // as we have used &User while defining u
	fmt.Println(u.Name)
}

func Modify(u *User) {
	*u = User{Name: "Bob"}
}


// More examples with int
package main

import (
	"fmt"
)

func testPtr(s *int) {

	fmt.Println("Value in func:")

	fmt.Printf("\ts = %v\n", s)
	fmt.Printf("\t&s = %v\n", &s)
	fmt.Printf("\t*s = %v\n", *s)
	// to change actually value under s pointer
	// we need to dereference s pointer like s to
	// to change underlying data, because s here
	// is just a copy of pointer (s = &tmp <== won't work)
	tmp := 11
	(*s) = tmp

	fmt.Println("Upd in func:")
	fmt.Printf("\t*s = %v\n", *s)
	fmt.Printf("\ts = %v\n", s)

}

type test struct {
	val int
}

func main() {
	var s int
	s = 10

	fmt.Println("Init:")
	fmt.Printf("\t&s = %v\n", &s)
	fmt.Printf("\ts = %v\n", s)

	testPtr(&s)

	fmt.Println("After Func:")
	fmt.Printf("\t&s = %v\n", &s)
	fmt.Printf("\ts = %v\n", s)

}
