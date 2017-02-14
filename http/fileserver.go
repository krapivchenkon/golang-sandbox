package main

import (
	"net/http"
)

// Fileserver for serving static files
func main() {
	http.ListenAndServe(":8888", http.FileServer(http.Dir(".")))
}
