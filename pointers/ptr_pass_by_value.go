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
