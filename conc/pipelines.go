package main

// https://blog.golang.org/pipelines
import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
)

func gen(N int) chan int {
	out := make(chan int)
	if N > 0 {

		go func() {
			for i := 0; i < N; i++ {
				out <- i
			}
			close(out)
		}()
	}

	return out
}

type Msg struct {
	ch  string
	pld int
}

func sq(ch string, in chan int) chan Msg {
	out := make(chan Msg)
	go func() {
		for v := range in {
			out <- Msg{ch, v * v}
		}
		close(out)
	}()

	return out
}

// fan-out, fan-in
func merge(cs ...chan Msg) chan Msg {
	out := make(chan Msg)
	var wg sync.WaitGroup

	// function that will be spawned in separate goroutine to read from one of input channels
	_merge := func(in chan Msg) {
		for m := range in {
			out <- m
		}
		wg.Done()
	}

	// initialilze semaphore
	wg.Add(len(cs))

	for _, v := range cs {
		go _merge(v)
	}

	//start separate goroutine to close out channel after all read is done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Explicit Cancellation functions
func genCncl(done <-chan struct{}, N int) chan int {
	out := make(chan int)
	if N > 0 {

		go func() {
			defer fmt.Printf("Channel %v closed\n", "GEN")
			defer close(out)
			for i := 0; i < N; i++ {
				select {
				case out <- i:
					fmt.Printf("Writing value %v\n", i)
				case <-done:
					fmt.Printf("Channel %v closed via done\n", "gen")
					return
				}
			}
		}()
	}

	return out
}

func sqCncl(done <-chan struct{}, ch string, in chan int) chan Msg {
	out := make(chan Msg)
	go func() {
		defer fmt.Printf("Channel square %v closed\n", ch)
		defer close(out)
		for v := range in {
			select {
			case out <- Msg{ch, v * v}:

			case <-done:
				// if closed {
				// 	fmt.Printf("Channel %v closed\n", "square")
				// }
				fmt.Printf("Channel %v closed via done\n", "square")
				return
			}
		}
	}()

	return out
}
func mergeCancel(done <-chan struct{}, cs ...chan Msg) chan Msg {
	out := make(chan Msg)
	var wg sync.WaitGroup
	// defer close(out)

	// function that will be spawned in separate goroutine to read from one of input channels
	_merge := func(ch string, in chan Msg) {
		defer fmt.Printf("Channel _merge %v closed\n", ch)
		defer wg.Done()

		for m := range in {
			select {
			case out <- m:
				fmt.Printf("Send out from %v \n", ch)
			case <-done:
				// if closed {
				// 	fmt.Printf("Channel %v closed due to done closed\n", ch)
				// }
				fmt.Printf("Channel %v closed via done\n", ch)
				return
			}
		}
	}

	// initialilze semaphore
	wg.Add(len(cs))

	for i, v := range cs {
		go _merge(fmt.Sprintf("c%v", i+1), v)
	}

	//start separate goroutine to close out channel after all read is done
	go func() {
		// defer close(out)
		fmt.Println("Start waiting")
		wg.Wait()
		fmt.Println("All finished")
		close(out)
	}()

	return out

}

func main() {
	var command string
	flag.StringVar(&command, "test", "default", "Enter test number")

	flag.Parse()
	switch command {
	case "1": // run -test 1
		for n := range sq("c1", gen(5)) {
			fmt.Printf("msg:%v\n", n)
		}
	case "fan-out": // run -test fan-out
		gen_num := gen(10)
		c1 := sq("c1", gen_num)
		c2 := sq("c2", gen_num)
		for n := range merge(c1, c2) {

			fmt.Printf("el:%v\n", n)
		}
	case "cancellation":
		fmt.Println("Explicit cancellation")
		// as we are using unbeffered channels if out channel will fail to read - all sending channels will be blocked
		// Goroutines are not garbage collected; they must exit on their own.

		done := make(chan struct{})
		defer fmt.Println("Closing main")
		// defer close(done)
		in := genCncl(done, 10)

		// Distribute the sq work across two goroutines that both read from in.
		c1 := sqCncl(done, "c1", in)
		c2 := sqCncl(done, "c2", in)

		// Consume the first value from output.
		out := mergeCancel(done, c1, c2)
		fmt.Println(<-out) // 4 or 9
		fmt.Println(<-out) // 4 or 9
		// for v := range out {
		// 	fmt.Println(v)
		// }

		// L:
		// 	for {
		// 		select {
		// 		case v, ok := <-out:
		// 			if !ok {
		// 				fmt.Println("Merged channel closed")
		// 				break L
		// 			}
		// 			fmt.Println(v)
		// 		}
		// 	}

		// send broadcast event to all subscribed gourutines
		close(done)
		v, ok := <-out
		if !ok {
			fmt.Println("Merged channel closed ")

		}
		fmt.Println(v)

	default:

	}

	bufio.NewReader(os.Stdin).ReadString('\n')
	// var inp string
	// fmt.Scan(&inp)
	fmt.Println("Exit...")
	// close()

}
