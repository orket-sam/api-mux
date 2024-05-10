package main

import (
	"encoding/json"
	"net/http"
)

func (s *APIServer) HandlerAccounts(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewServer(listenAddress string) *APIServer {
	return &APIServer{listenAddress}
}

func WriteJson(w http.ResponseWriter, v any, status int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func MakeHttpHandler(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// log.Fatal(err.Error())
			println("hdj")
		}
	}
}
