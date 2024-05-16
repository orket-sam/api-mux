package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=orket dbname=go_bank password=mysecretpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())

	}
	log.Println("Db connected succesfully")

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())

	}

	log.Println("Db pinged succesfully")

	// defer db.Close()

	return &PostgresStore{db}, err

}

func (s *PostgresStore) CreateAccount(account *Account) error {
	query := fmt.Sprintf("insert into account(first_name,last_name,number)values('%s' ,'%s','%s')", account.FirstName, account.LastName, account.Number)
	_, err := s.DB.Exec(query)
	return err
}

func (s *PostgresStore) UpdateAccount(account *Account) error {
	query := "update account set first_name=$1,last_name=$2 where id=$3 "

	_, err := s.DB.Query(query, account.FirstName, account.LastName, account.ID)

	return err
}
func (s *PostgresStore) DeleteAccount(id int) error {
	query := "delete from account where id=$1"
	_, err := s.DB.Query(query, id)

	return err
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	var account Account
	query := "select * from account where id=$1"
	err := s.DB.QueryRow(query, id).Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.Created_AT)
	if err != nil {
		log.Println("lol" + err.Error())
	}

	return &account, err
}
