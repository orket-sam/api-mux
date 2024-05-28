package main

import (
	"fmt"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
)

func HandlerCheckCreateAccReqOnCreate(accrequest *CreateAccountRequest, w http.ResponseWriter) error {
	if accrequest.FirstName == "" || accrequest.LastName == "" {
		apiError := APIError{"first_name and last_name required"}
		http.Error(w, apiError.Message, http.StatusBadRequest)
		return &apiError
	}
	return nil
}

func WithJwt(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-token")
		f(w, r)
	}
}

func ValidateJwt(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("secret")
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			println("Invalid signing")

			return nil, fmt.Errorf("invalid signing method used")

		}
		println("valid signing")
		return []byte(secret), nil
	})
}
