package main

import (
	"fmt"
	//  "os"
)

type msg struct {
	from int
	s    int
}

func produce(dataChannel chan msg, id int) {
	for i := 0; i < 10; i++ {
		//fmt.Fprintf(os.Stdout, "[%v] - %-3d\n", id, i)
		dataChannel <- msg{from: id, s: i}

	}
}

func main() {
	dataChannel := make(chan msg)

	go produce(dataChannel, 1)
	go produce(dataChannel, 2)
	go produce(dataChannel, 3)

	for i := 0; i < 30; i++ {
		data := <-dataChannel
		fmt.Println(data)
	}
}
