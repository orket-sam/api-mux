package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDb() error {
	connStr := "user=orket dbname=go-bank sslmode=disable password=mysecret"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)

	} else {
		log.Println("db connected successfully")
	}
	query := "insert into account (first_name,last_name,account_number)values($1,$2,$3)"
	if _, err := db.Query(query, "sam", "orket", "48728977438"); err != nil {
		println(err.Error())
	}
	return nil
}
