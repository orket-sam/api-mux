package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) AccountHandler(w http.ResponseWriter, r *http.Request) error {
	return s.CreateAccountHandler(w, r)
}

func (s *APIServer) CreateAccountHandler(w http.ResponseWriter, r *http.Request) error {

	var newAccount Account
	if len(newAccount.FirstName) == 0 {
		return WriteJson(w, "first_name and last_name required", 422)
	}
	json.NewDecoder(r.Body).Decode(&newAccount)
	fmt.Println(newAccount)
	return nil
}

func WriteJson(w http.ResponseWriter, v any, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

func MakeHttpHandler(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(err.Error())
		}
	}
}

func (s *APIServer) RunServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", MakeHttpHandler(s.AccountHandler))
	http.ListenAndServe(s.ListenAddress, router)
}
