package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

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
