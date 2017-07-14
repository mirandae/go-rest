package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	defaultPortNum = 8080
	defaultAddr    = "localhost"
)

func main() {

	http.HandleFunc("/", GET)
	http.HandleFunc("/get", GET)
	http.HandleFunc("/put", PUT)
	http.HandleFunc("/post", POST)
	http.HandleFunc("/delete", DELETE)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", defaultAddr, defaultPortNum), nil))
}
