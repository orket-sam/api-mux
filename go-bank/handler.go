package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
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

		tokenString := r.Header.Get("jwt")
		print(tokenString)
		token, err := ValidateJwt(tokenString)
		if err != nil {
			WriteJson(w, err.Error(), http.StatusForbidden)
			return
		}
		claimId := token.Claims.(jwt.MapClaims)["id"]
		id, _ := GetID(r)
		if claimId != id {
			WriteJson(w, "No access", http.StatusForbidden)
			return
		}
		f(w, r)
	}

}

func ValidateJwt(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("secret")

	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s ", t.Header["alg"])

		}

		return []byte(secret), nil
	})
}
