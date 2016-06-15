package main

import (
	"fmt"
	"net/http"
)

func restHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go runner in minimal Docker container")
}

func main() {

	http.HandleFunc("/", restHandler)

	fmt.Println("Started, serving at 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
