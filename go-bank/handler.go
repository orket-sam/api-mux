package main

import "net/http"

func HandlerCheckCreateAccReqOnCreate(accrequest *CreateAccountRequest, w http.ResponseWriter) error {
	if accrequest.FirstName == "" || accrequest.LastName == "" {
		apiError := APIError{"first_name and last_name required"}
		http.Error(w, apiError.Message, http.StatusBadRequest)
		return &apiError
	}
	return nil
}
