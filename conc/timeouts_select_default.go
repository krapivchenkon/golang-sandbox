package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var README string = `
Taken from https://blog.golang.org/go-concurrency-patterns-timing-out-and
in this example you should test different combinations of buffered(1) and unbuffered channel in Query func
ch := make(chan Result)
 and ch := make(chan Result,1)

Removing default select case for Query func
- in case of unbuffered channel on of the goroutines will hang forever

second query here won't be cancelled but result from first one will be available
`

type Conn struct {
	id string
}

type Result struct {
	msg     string
	timeout string
}

func (c *Conn) DoQuery(query string) Result {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	dur := r1.Float64()*3.0 + 1

	log.Print(fmt.Sprintf("channel %v dur %.2fs", c.id, dur))
	<-time.After(time.Duration(dur) * time.Second)
	return Result{msg: fmt.Sprintf("[resp for \"%v\" from %v]", query, c.id), timeout: fmt.Sprintf("%.2fs", dur)}
}

func Query(query string, conns ...Conn) <-chan Result {
	ch := make(chan Result, 1)
	for _, conn := range conns {
		go func(c Conn) {
			defer log.Printf("goroutine %v exited", c.id)
			select {
			case ch <- c.DoQuery(query):
				log.Printf("Response arrived from conn %v", c.id)
			default:
				log.Printf("Default select for conn %v", c.id)
			}
		}(conn)
	}
	// <-time.After(4 * time.Second)

	return ch
}

func main() {
	c1 := Conn{id: "c1"}
	c2 := Conn{id: "c2"}
	out := Query("test", c1, c2)
	log.Print(<-out)

	bufio.NewReader(os.Stdin).ReadString('\n')
}
