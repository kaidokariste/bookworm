package main

import (
	"net/http"
	"log"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("/public"))
	mux.Handle("/", fs)
	log.Println("Listening...")
	http.ListenAndServe(":8080", fs)
}
