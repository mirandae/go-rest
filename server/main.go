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

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/notes", NotesListPage)
	http.HandleFunc("/notes/", NotePage)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", defaultAddr, defaultPortNum), nil))
}
