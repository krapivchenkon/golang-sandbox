// Interface project main.go
// goodread http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := new(person)
	fmt.Println(p.talk())
	fmt.Println(p.walk())

	o := new(policeOfficer)
	fmt.Println(o.talk())
	fmt.Println(o.walk())
	o.badgeNumber = 1000
	fmt.Println(o.badgeNumber)
	fmt.Println(o.run())

	fmt.Println(run(15))

	// fmt.Println("Check type:",o.(type))
	var r = reflect.TypeOf(o)
    fmt.Printf("Other:%v\n", r)             
	// switch t := o.(type) {
 //    case int:
 //        fmt.Printf("Integer: %v\n", t)
 //    case float64:
 //        fmt.Printf("Float64: %v\n", t)
 //    case string:
 //        fmt.Printf("String: %v\n", t)
 //    case bool:
 //        fmt.Printf("Bool: %v\n", t)
 //    case []interface {}:
 //        for i,n := range t {
 //            fmt.Printf("Item: %v= %v\n", i, n)
 //        }
 //    default:
 //        var r = reflect.TypeOf(t)
 //        fmt.Printf("Other:%v\n", r)             
 //    }
}

type person struct {
	firstname string
	lastname  string
}

type policeOfficer struct {
	badgeNumber int
	precinct    string
}

type behaviors interface {
	talk() string
	walk() int
	run() int
}

//person implementation
func (p *person) talk() string {
	return "hi there!"
}

func (p *person) walk() int {
	return 10
}

//officier implementation
func (o *policeOfficer) talk() string {
	return "Have a good day"
}

func (o *policeOfficer) walk() int {
	return 20
}

//func [param list] [interface func name] [interface func return type]
func (o *policeOfficer) run() int {
	return 50
}

//regular function
//func [name] [param list] [return type]
func run(s int) int {
	return s
}
