package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) AccountHandler(w http.ResponseWriter, r *http.Request) error {
	return WriteJson(w, "good morning", 200)
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
