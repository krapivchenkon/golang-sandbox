package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}
	response, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		// _, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(string(body[:100]))
	}
}

// TODO: client reading and sending json sample

// TODO: Client using Timeout settings

// TODO: Client sends form object

// TODO: client http/https select
