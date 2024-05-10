package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresStore() (*PostgresStore, error) {

	connStr := "user=orket dbname=go_bank sslmode=disable password=mysecretpassword host=127.0.0.1 port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())

	}
	println("Database pinged successfully")

	return &PostgresStore{db}, err

}
