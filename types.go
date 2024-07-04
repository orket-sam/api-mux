package main

import (
	"database/sql"
	"net/http"
	"time"
)

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIServer struct {
	ListenAddress string
	Db            Storage
}

type APIError struct {
	Message string
}

type Storage interface {
	CreateAccount(Account) error
}

type PostgresStorage struct {
	Db *sql.DB
}

func NewAPIServer(listenAddress string, db Storage) *APIServer {

	return &APIServer{listenAddress, db}
}

type Account struct {
	Id            int32     `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	CreatedAT     time.Time `json:"created_at"`
}

func (e *APIError) Error() string {
	return "Error: " + e.Message
}
