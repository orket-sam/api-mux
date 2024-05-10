package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) HandlerAccounts(w http.ResponseWriter, r *http.Request) error {
	return s.Storage.CreateAccount(&Account{})
}

func NewServer(listenAddress string, store Store) *APIServer {
	return &APIServer{listenAddress, store}
}

func WriteJson(w http.ResponseWriter, v any, status int) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func MakeHttpHandler(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			WriteJson(w, err.Error(), 404)
		}
	}
}

func (s *APIServer) RunsServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", MakeHttpHandler(s.HandlerAccounts))

	err := http.ListenAndServe(s.ListenAddress, r)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("server is up and running")
}
