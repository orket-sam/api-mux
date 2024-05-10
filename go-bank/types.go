package main

import (
	"database/sql"
	"net/http"
	"time"
)

type APIServer struct {
	ListenAddress string
	Storage       Store
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type Account struct {
	Created_AT time.Time `json:"created_at"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Number     string    `json:"account_number"`
	Balance    int       `json:"balance"`
	ID         int       `json:"id"`
}

type Store interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	DB *sql.DB
}
