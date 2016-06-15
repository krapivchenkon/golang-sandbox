// WebServer project main.go
package main

import (
	"fmt"
	"html"
	// "io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8181", nil)

	// resp, err := http.Get("http://www.google.com")
	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
