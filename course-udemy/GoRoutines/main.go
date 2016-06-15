// GoRoutines project main.go
package main

import (
	"fmt"
)

// can be done with waiting group as well

func main() {
	cars := fillCars()
	msg := make(chan string)
	go showCars(msg, cars, "[1] goroutine")
	go showCars(msg, cars, "[2] goroutine")
	go showCars(msg, cars, "[0] goroutine")

	go func(msgstr string) {
		msg <- msgstr
		// fmt.Println()
	}("going")

	//time.Sleep(2 * time.Second)

	// var input string
	// fmt.Scanln(&input)
	fmt.Println("done")
	// this is actually should be the number of messages that we expect on the channel
	// this way is good when we know the number of gourutines
	goroutines_spawned := 16
	for i := 0; i < goroutines_spawned; i++ {
		select {
		case mss := <-msg:
			fmt.Println(mss)
		}
	}

}

type Cars map[string]int

func fillCars() map[string]int {
	cars := make(map[string]int)
	cars["Jeep"] = 23
	cars["Buick"] = 11
	cars["Ford"] = 15
	cars["Chevy"] = 9
	cars["Nissan"] = 29
	return cars
}

func showCars(msg chan string, c Cars, m string) {
	for key, i := range c {
		msg <- fmt.Sprintf("[%[1]v] %[2]v = %[3]v %[4]v", m, i, key, c[key])
	}
}
