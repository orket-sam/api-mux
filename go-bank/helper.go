package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (apiError *APIError) Error() string {

	return apiError.Message
}

func GetID(r *http.Request) (int, error) {
	idString := mux.Vars(r)["id"]
	if id, err := strconv.Atoi(idString); err != nil {
		return id, err
	} else {
		return id, nil
	}
}
