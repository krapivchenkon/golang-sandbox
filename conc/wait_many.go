package main

import (
	"fmt"
	"time"
)

func WaitMany(a, b chan bool) {
	for a != nil || b != nil {
		select {
		case <-a:
			// based on the fact that nil channel always blocks either for receiving or for sending
			// but closed channel is always ready to receive
			a = nil
		case <-b:
			b = nil
		}
	}
}

func main() {
	a, b := make(chan bool), make(chan bool)
	t0 := time.Now()
	go func() {
		close(a)
		close(b)
	}()
	WaitMany(a, b)
	fmt.Printf("waited %v for WaitMany\n", time.Since(t0))
}
