package main

import "fmt"

type stack []int

func (s stack) Empty() bool { return len(s) == 0 }
func (s stack) Head() int   { return s[len(s)-1] }
func (s *stack) Put(i int)  { (*s) = append((*s), i) }
func (s *stack) Pop() int {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

// TODO: rewrite stack implementation in order  remove allocations on every add

func main() {
	var s stack

	for i := 0; i < 3; i++ {
		s.Put(i)
		fmt.Printf("len=%d\n", len(s))
		fmt.Printf("peek=%d\n", s.Head())
	}

	for !s.Empty() {
		i := s.Pop()
		fmt.Printf("len=%d\n", len(s))
		fmt.Printf("pop=%d\n", i)
	}
}
