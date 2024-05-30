package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func (s *APIServer) HandlerAccounts(w http.ResponseWriter, r *http.Request) error {

	idString := mux.Vars(r)["id"]

	switch method := r.Method; method {
	case "POST":
		return s.HandlerCreateAccount(w, r)
	case "DELETE":
		return s.HandlerDeleteAccount(w, r)
	case "PUT":
		return s.HandlerUpdateAccount(w, r)

	case "GET":
		if idString != "" {
			return s.HandlerGetAccountById(w, r)

		}
		return s.HandlerGetAllAccounts(w, r)
	default:
		return WriteJson(w, "invalid method", 500)
	}

}

func (s *APIServer) HandlerCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccReq := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(&createAccReq); err != nil {
		log.Println("parsing values error " + err.Error())
		http.Error(w, "parsing values error ", http.StatusBadRequest)

		return nil

	}
	if err := HandlerCheckCreateAccReqOnCreate(createAccReq, w); err != nil {
		return nil
	}

	account := NewAccount(createAccReq.FirstName, createAccReq.LastName)
	tokenString, err := CreateJwt(account)
	if err != nil {

		return WriteJson(w, APIError{err.Error()}, http.StatusBadRequest)
	}
	fmt.Println("JWT:  ", tokenString)
	if _, err := ValidateJwt(tokenString); err != nil {
		println("Invalid token")
	} else {
		println("************valid token*************")

	}
	if err := s.Storage.CreateAccount(account); err != nil {
		log.Println("create acc error: " + err.Error())
	}
	return WriteJson(w, createAccReq, 200)

}

func (s *APIServer) HandlerGetAllAccounts(w http.ResponseWriter, r *http.Request) error {

	accounts, err := s.Storage.GetAllAccounts()
	if err != nil {
		log.Println(err)
		return nil
	}
	return WriteJson(w, accounts, 200)

}

func (s *APIServer) HandlerUpdateAccount(w http.ResponseWriter, r *http.Request) error {

	idString := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idString)

	var accProfile *CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&accProfile); err != nil {
		log.Println("parsing values failed" + err.Error())
		http.Error(w, "wrong json format", http.StatusBadRequest)
		return nil
	}

	if err := s.Storage.UpdateAccount(accProfile, id); err != nil {
		log.Println("Update account error:" + err.Error())
		return nil
	}

	return WriteJson(w, &accProfile, 200)
}

func (s *APIServer) HandlerDeleteAccount(w http.ResponseWriter, r *http.Request) error {

	id, err := GetID(r)
	if err != nil {
		return err
	}

	if err := s.Storage.DeleteAccount(id); err != nil {
		return WriteJson(w, APIError{err.Error()}, http.StatusBadRequest)
	} else {
		return WriteJson(w, "Account deleted succesfully", http.StatusBadRequest)
	}

}

func (s *APIServer) HandlerGetAccountById(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(r)

	if err != nil {
		return err

	}
	if account, err := s.Storage.GetAccountByID(id); err != nil {
		// WriteJson(w, APIError{"account not found"}, http.StatusBadRequest)
		return err
	} else {
		return WriteJson(w, account, 200)
	}

}

func NewServer(listenAddress string, store Store) *APIServer {
	return &APIServer{listenAddress, store}
}

func WriteJson(w http.ResponseWriter, v any, status int) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func ReadJson(w http.ResponseWriter, r *http.Request, v any) error {

	return json.NewDecoder(r.Body).Decode(&v)
}

func MakeHttpHandler(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			WriteJson(w, err.Error(), 404)
		}
	}
}

func (s *APIServer) RunsServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", (MakeHttpHandler(s.HandlerAccounts)))
	r.HandleFunc("/{id}", WithJwt(MakeHttpHandler(s.HandlerAccounts)))

	err := http.ListenAndServe(s.ListenAddress, r)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("server is up and running")
}

func CreateJwt(account *Account) (string, error) {
	secret := os.Getenv("secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":           account.FirstName,
		"account_number": account.LastName,
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
