package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config interface {
	Get(key string) (string, error)
	Set(key, val string) error
}

var (
	Cfg InmemConfig
)

type InmemConfig struct {
	M map[string]string
}

type Responder struct {
	Cfg Config
}

func (c InmemConfig) Get(key string) (string, error) {
	val, ok := c.M[key]
	if ok {
		return val, nil
	} else {
		return "", errors.New("Tried to get a key which doesn't exist")
	}
}

func (c InmemConfig) Set(key, val string) error {
	c.M[key] = val
	return nil
}

func (r *Responder) Handler(w http.ResponseWriter, req *http.Request) {
	cfgOption, err := r.Cfg.Get("Option.Header")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("X-Config-Option", cfgOption)
	fmt.Fprintf(w, "This is the response body!")
}

func main() {
	responder := Responder{
		Cfg: InmemConfig{
			M: make(map[string]string),
		},
	}
	if err := responder.Cfg.Set("Option.Header", os.Args[1]); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", responder.Handler)
	http.ListenAndServe(":8080", nil)
}
