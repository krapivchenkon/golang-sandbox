package main

import "fmt"
import "time"

// Using "nil" Channels
// Send and receive operations on a nil channel block forver.
// This behavior can be used as a way to dynamically enable
// and disable case blocks in a select statement.

func main() {
	inch := make(chan int)
	outch := make(chan int)

	go func() {
		var in <-chan int = inch
		var out chan<- int
		var val int
		for {
			select {
			case out <- val:
				out = nil
				in = inch
			case val = <-in:
				out = outch
				in = nil
			}
		}
	}()

	go func() {
		for r := range outch {
			fmt.Println("result:", r)
		}
	}()

	time.Sleep(0)
	inch <- 1
	inch <- 2
	time.Sleep(3 * time.Second)
}
