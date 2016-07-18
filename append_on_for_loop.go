package main

import (
	"fmt"
)

// Example of how to deal when wee need to change slice while iterating

func main() {
	fmt.Println("Hello, playground")

	s := []int{1, 2, 3, 4}

	fmt.Println(len(s))

	cnt, cntRead := len(s), 0
	for i := 0; true; i++ {
		if cnt == cntRead {
			break
		}
		if s[i]%2 == 0 {
			l := []int{s[i] + 1, 101}
			s = append(s, l...)
			//s[cnt]=s[i]+1
			cnt += len(l)
		}
		cntRead++
		// fmt.Println(s[i])
	}
	fmt.Println(s)

}
