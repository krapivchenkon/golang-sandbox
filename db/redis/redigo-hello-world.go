package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Cannot connect")
		panic(err)
	}
	defer c.Close()

	// Setting simple keys as integers
	c.Do("SET", "k1", 1)
	c.Do("SET", "k2", "string value")
	n, _ := redis.Int(c.Do("GET", "k1"))
	fmt.Printf("%#v\n", n)
	n, _ = redis.Int(c.Do("INCR", "k1"))
	fmt.Printf("%#v\n", n)

	// test MGET
	reply, err := redis.Values(c.Do("MGET", "k1", "k2"))
	if err != nil {
		// handle error
	}
	fmt.Println(reply)
	// if _, err := redis.Scan(reply, &value1, &value2); err != nil {
	// handle error
	// }

	// Working with lists

	// Working with Hashes

	// Working with sets

	// Working with sorted sets

}
