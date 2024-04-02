package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// parse netowork address flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register handler functions
	// this is where all of your routes (URL patterns) will go
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Start server on TCP network address ":4000"
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
