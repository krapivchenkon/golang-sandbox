// Generated private key
//      $ openssl genrsa -out server.key 2048
// To generate a certificate
//      $ openssl req -new -x509 -key server.key -out server.pem -days 3650
// Install certificate on your system
// READ: https://blog.bracelab.com/achieving-perfect-ssl-labs-score-with-go
//       https://gist.github.com/denji/12b3a568f092ab951456
package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":443", "server.pem", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
