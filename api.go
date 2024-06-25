package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, v any) error {
	return json.NewEncoder(w).Encode(v)
}

func (s *APIServer) Runserver() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		WriteJson(w, "Ciao, welcome to orket's api")
	})
	http.ListenAndServe(s.ListenAddress, r)
}
