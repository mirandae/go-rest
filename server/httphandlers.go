package main

import (
	"fmt"
	"html"
	"net/http"
	"projects/go-rest/server/store"
)

// global store
var notesStore = store.NoteStoreImpl{}

// HomePage picks the function to handle a request to
// the root ('/') page
func HomePage(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// Serve the resource.
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
	default:
		fmt.Fprintf(w, "Error: Invalid request")
	}

}

// NotesListPage picks the function to handle a request to
// the 'notes' page
func NotesListPage(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// return list of valid note IDs

	default:
		// Give an error message.
	}

}

// NotePage picks the function to handle a request to
// a specific note
func NotePage(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// return the note corresponding to the ID
	case "POST":
		// Create a new note, overwriting the old one, if it existed
	case "PUT":
		// Update the note.
	case "DELETE":
		// Delete the note
	default:
		// Give an error message.
	}
}
