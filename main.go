package main

import (
	"log"
	"net/http"
	"projects/go-rest/server"
)

func main() {

	http.HandleFunc("/", nil)
	http.HandleFunc("/get", server.GET)
	http.HandleFunc("/put", server.PUT)
	http.HandleFunc("/post", server.POST)
	http.HandleFunc("/delete", server.DELETE)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
