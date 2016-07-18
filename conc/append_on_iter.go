package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, playground")

	s := []int{1, 2, 3, 4}

	//  s := make([]int, 10)
	//  s[0], s[1], s[2], s[3] = 1, 2, 3, 4
	//  fmt.Println(len(s))

	cnt := len(s)

	ch := make(chan int, len(s))
	resp := make(chan int, 2)

	go func() {
		defer close(ch)
		for _, v := range s {
			ch <- v
		}
	L:
		for {
			select {
			case v, ok := <-resp:
				if !ok {
					fmt.Println("closed resp")

					break L
				} else {

					ch <- v
				}
			}
		}

	}()

	cntRead := 0
	for v := range ch {
		cntRead++

		if v%2 == 0 {

			resp <- v + 1
			resp <- v + 10
			cnt++

			//fmt.Println(s)
		}
		fmt.Println(v)
		if cnt == cntRead {
			close(resp)
			break
		}
	}

	<-time.After(1 * time.Second)

}
