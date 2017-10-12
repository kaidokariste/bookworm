package main

import (

	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Page struct {
	Title string
}

var templates = template.Must(template.ParseFiles("src/eugeii/header.html"))

func RootHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}
	templates.ExecuteTemplate(response, "header.html", Page{Title: "Home"})
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	_, error := os.Stat(filepath.Join(".", "src/eugeii/static/css", "bootstrap.min.css")) // test when it reaches to path
	log.Print(error)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("src/eugeii/static/"))))

	addr := fmt.Sprintf("%s:%d", "localhost", 8080)
	log.Printf("Listening on %s", addr)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}