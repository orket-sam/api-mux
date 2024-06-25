package main

import "net/http"

type ApiError struct {
	Message string
}

type APIFunc interface {
	func(http.ResponseWriter, *http.Request) error
}

type APIServer struct {
	ListenAddress string
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{listenAddress}
}
