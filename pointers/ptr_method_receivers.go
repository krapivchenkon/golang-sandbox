package main

import "fmt"

// Method receivers are like regular function arguments.
// If it's declared to be a value then your function/method
// gets a copy of your receiver argument. This means making
// changes to the receiver will not affect the original
// value unless your receiver is a map or slice variable
// and you are updating the items in the collection or the
// fields you are updating in the receiver are pointers.

type data struct {
	num   int
	key   *string
	items map[string]bool
}

func (this *data) pmethod() {
	this.num = 7
}

func (this data) vmethod() {
	this.num = 8
	*this.key = "v.key"
	this.items["vmethod"] = true
}

func main() {
	key := "key.1"
	d := data{1, &key, make(map[string]bool)}

	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=1 key=key.1 items=map[]

	d.pmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=7 key=key.1 items=map[]

	d.vmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=7 key=v.key items=map[vmethod:true]
}
