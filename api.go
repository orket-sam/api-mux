package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	ListenAddress string
}

type APIFunc func(http.ResponseWriter, *http.Request) error

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{ListenAddress: listenAddress}
}

func WriteJson(v any, w http.ResponseWriter, statusCode int) error {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func MakeHttpHandler(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(err.Error(), w, http.StatusBadRequest)
		}

	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", MakeHttpHandler(s.HandleAccount))
	log.Println("Api running on port: ", s.ListenAddress)
	http.ListenAndServe(s.ListenAddress, router)
}

func (s *APIServer) HandleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {

		return s.HandleGetAccount(r, w)
	}
	if r.Method == "POST" {
		return s.HandleCreateAccount(r, w)

	}
	if r.Method == "DELETE" {
		return s.HandleDeleteAccount(r, w)

	}
	return fmt.Errorf(r.Method, "method not allowed")
}

func (s *APIServer) HandleCreateAccount(r *http.Request, w http.ResponseWriter) error {
	return nil
}
func (s *APIServer) HandleDeleteAccount(r *http.Request, w http.ResponseWriter) error {
	return nil
}
func (s *APIServer) HandleGetAccount(r *http.Request, w http.ResponseWriter) error {
	return WriteJson("welcome to api version 1", w, http.StatusOK)
}
