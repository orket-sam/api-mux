package main

import "net/http"

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIServer struct {
	ListenAddress string
}

type APIError struct {
	Message string
}

func NewAPIServer(listenAddress string) *APIServer {

	return &APIServer{listenAddress}
}

type Account struct {
	FirstName string
	lastName  string
	//  floa

}
