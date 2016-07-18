package main

import "fmt"

func main() {
	ch := make(chan bool, 2)
	fmt.Printf("chan cap:%v\n", cap(ch))
	ch <- true
	ch <- true
	close(ch)
	for i := 0; i < cap(ch)+1; i++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}
}
