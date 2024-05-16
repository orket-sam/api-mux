package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) HandlerAccounts(w http.ResponseWriter, r *http.Request) error {

	switch method := r.Method; method {
	case "POST":
		return s.HandlerCreateAccount(w, r)
	case "DELETE":
		return s.HandlerDeleteAccount(w, r)
	case "PUT":
		return s.HandlerUpdateAccount(w, r)
	case "GET":
		return s.HandlerGetAccountById(w, r)
	default:
		return WriteJson(w, "invalid method", 500)
	}

}

func (s *APIServer) HandlerCreateAccount(w http.ResponseWriter, r *http.Request) error {
	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Fatal("parsing values failed" + err.Error())
	}
	log.Println(account)
	return s.Storage.CreateAccount(&account)

}

func (s *APIServer) HandlerUpdateAccount(w http.ResponseWriter, r *http.Request) error {

	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Fatal("parsing values failed" + err.Error())
	}
	log.Println("feel good")

	return s.Storage.UpdateAccount(&account)
}

func (s *APIServer) HandlerDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)

	return s.Storage.DeleteAccount(id)
}

func (s *APIServer) HandlerGetAccountById(w http.ResponseWriter, r *http.Request) error {
	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)
	account, _ := s.Storage.GetAccountByID(id)
	return WriteJson(w, account, 200)

}

func NewServer(listenAddress string, store Store) *APIServer {
	return &APIServer{listenAddress, store}
}

func WriteJson(w http.ResponseWriter, v any, status int) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func ReadJson(w http.ResponseWriter, r *http.Request, v any) error {

	return json.NewDecoder(r.Body).Decode(&v)
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
	r.HandleFunc("/{id}", MakeHttpHandler(s.HandlerAccounts))

	err := http.ListenAndServe(s.ListenAddress, r)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("server is up and running")
}
