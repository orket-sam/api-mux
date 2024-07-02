package main

import "net/http"

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIServer struct {
	ListenAddress string
	// Store         Storage
}

type APIError struct {
	Message string
}

type Storage interface {
	CreateAccount(Account) error
}

func NewAPIServer(listenAddress string) *APIServer {

	return &APIServer{listenAddress}
}

type Account struct {
	Id            int32   `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
}
