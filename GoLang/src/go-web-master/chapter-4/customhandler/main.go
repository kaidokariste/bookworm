package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct { // struct named messageHandler is created
	message string
}

//method with signature ServeHTTP is implemented to make type messageHandler
// implement handler
func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()

	mh1 := &messageHandler{"Welcome to Go Web Development"} //instance created
	mux.Handle("/welcome", mh1) // ServeMux.Handle is called to register handlers

	mh2 := &messageHandler{"net/http is awesome"}
	mux.Handle("/message", mh2)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
