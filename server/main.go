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

	http.HandleFunc("/", server.GET)
	http.HandleFunc("/get", server.GET)
	http.HandleFunc("/put", server.PUT)
	http.HandleFunc("/post", server.POST)
	http.HandleFunc("/delete", server.DELETE)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", defaultAddr, defaultPortNum), nil))
}
