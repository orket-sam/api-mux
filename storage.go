package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDb() (*PostgresStorage, error) {
	connStr := "user=orket dbname=go-bank sslmode=disable password=mysecret"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println(err.Error())
		return nil, err

	} else {
		log.Println("db connected successfully")
		return &PostgresStorage{db}, nil
	}

}

func (storage *PostgresStorage) CreateAccount(account Account) error {
	query := "insert into account(first_name,last_name,account_number)values($1,$2,$3)"
	sqlResult, err := storage.Db.Exec(query, account.FirstName, account.LastName, account.AccountNumber)
	if err != nil {
		return &APIError{err.Error()}
	} else {
		fmt.Println(sqlResult.RowsAffected())
	}

	return nil
}
