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
	// defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())

	}
	println("Database pinged successfully")

	return &PostgresStore{db}, err

}

func (s *PostgresStore) CreateAccount(*Account) error {

	query := `insert into account(first_name,last_name,number,balance)values('shanaya','wangari',12536819,129089)`
	_, err := s.DB.Query(query)
	return err
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
