package main

import (
	"fmt"
	"html"
	"net/http"
)

// GET handler
func GET(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
}

// PUT handler
func PUT(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
}

// POST handler
func POST(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
}

// DELETE handler
func DELETE(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
}
