package main

import (
	"fmt"
	"log"
	"net/http"
)

// if the function has signature func(http.ResponseWriter, *http.Request)
// it can be converted to HandleFunc
func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}

func main() {
	mux := http.NewServeMux()

	// Convert the messageHandler function to a HandlerFunc type
	mh := http.HandlerFunc(messageHandler)
	// syntax func (mux *ServeMux) Handle(pattern string, handler Handler)
	// mh - the HandlerFunc instance is created and added to mux.Handle to handle
	// requests coming from /welcome page
	mux.Handle("/welcome", mh)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
