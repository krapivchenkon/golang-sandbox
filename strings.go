package main

import "fmt"

func main() {
	fmt.Println(string("Hello"[1]))             // ASCII only
	fmt.Println(string([]rune("Hello, 世界")[1])) // UTF-8
	fmt.Println(string([]rune("Hello, 世界")[8])) // UTF-8
	for i, v := range "Hello, 世界" {
		// fmt.Println("test")
		fmt.Printf("char: %c at pos: %d \n", v, i)
	}
	// convert string to represent characters array
	chars := []rune("Hello, 世界")
	for i, v := range chars {
		fmt.Printf("char: %c at pos: %d \n", v, i)
	}

	fmt.Printf("%c\n", chars[7])

	str_from_runes := string([]rune{'h', 'e', 'l', 'l', 'o', '☃'})

	fmt.Println(str_from_runes)
}
