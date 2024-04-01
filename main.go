package main

import (
	"log"
	"net/http"
)

// Home handler function
func home(w http.ResponseWriter, r *http.Request) {
  // if the current URL path does not exactly match "/" then it is not the home page
  // return a 404
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()

	// Register handler functions
	// this is where all of your routes (URL patterns) will go
	mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("/snippet/create", snippetCreate)

	// Start server on TCP network address ":4000"
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
