package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "communalAPI/config"
	. "communalAPI/dao"

)

var config = Config{}
var dao = PaymentDAO{}

// GET list of movies
func AllPaymentsEndPoint(w http.ResponseWriter, r *http.Request) {
	payment, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, payment)
}

// // GET a movie by its ID
func FindPaymentEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payment, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, payment)
}


func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.DBUrl = config.DBUrl
	dao.AuthDatabase = config.AuthDatabase
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/payments", AllPaymentsEndPoint).Methods("GET")
	/*r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")*/
	r.HandleFunc("/payments/{id}", FindPaymentEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}
}