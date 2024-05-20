package main

import (
	"fmt"
	"net/http"
)

func HandlerCheckCreateAccReqOnCreate(accrequest *CreateAccountRequest, w http.ResponseWriter) error {
	if accrequest.FirstName == "" || accrequest.LastName == "" {
		apiError := APIError{"first_name and last_name required"}
		http.Error(w, apiError.Message, http.StatusBadRequest)
		return &apiError
	}
	return nil
}

func WithJWT(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware is working")
		handlerFunc(w, r)
	}
}
