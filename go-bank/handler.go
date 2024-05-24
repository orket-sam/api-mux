package main

import (
	"fmt"
	"log"
	"net/http"

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

func WithJWT(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware is working")
		tokenString := r.Header.Get("Authorization")
		_, err := ValidateJwt(tokenString)
		if err != nil {
			log.Println(err)
		}
		handlerFunc(w, r)
	}
}

func ValidateJwt(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("my_secret_key"), nil
	})

}
