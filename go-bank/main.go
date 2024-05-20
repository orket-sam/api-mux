package main

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func main() {
	// store, _ := NewPostgresStore()
	// server := NewServer(":2000", store)
	// server.RunsServer()

}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_AUTH_SECRET")
	jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method(*jwt.SigningMethodHMAC); !ok {
			// return nil,fmt.Errorf("unexpected  signing me")
		}

		return nil, nil
	})

	return nil, nil
}
