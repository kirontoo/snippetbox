package main

import "net/http"

// returns a servermux containing our application routes
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// flow: 
	// logRequest <-> secureHeaders <-> servemux <-> application handler
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
