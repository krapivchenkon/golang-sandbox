package main

// from https://www.reddit.com/r/golang/comments/3vndco/why_no_tail_call_elimination/
// about tail recursion in go https://www.goinggo.net/2013/09/recursion-and-tail-calls-in-go_26.html
import "fmt"

type facfunc func(n, acc int) (facfunc, int, int)

func fac(n, acc int) (facfunc, int, int) {
	if n < 1 {
		return nil, 0, acc
	}

	return fac, n - 1, acc * n
}

func Factorial(n int) int {
	f := facfunc(fac)
	acc := 1

	// Trampoline the tail calls
	for f != nil {
		f, n, acc = f(n, acc)
	}
	return acc
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i, Factorial(i))
	}
}
