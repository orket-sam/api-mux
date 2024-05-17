package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

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
	// query := fmt.Sprintf("insert into account(first_name,last_name,number)values('%s' ,'%s','%s')", account.FirstName, account.LastName, account.Number)
	query := "insert into account (first_name,last_name,number)values($1,$2,$3)"
	_, err := s.DB.Exec(query, account.FirstName, account.LastName, account.Number)
	return err
}

func (s *PostgresStore) UpdateAccount(updateAcc *CreateAccountRequest, id int) error {
	query := "update account set first_name=$1,last_name=$2 where id=$3 "

	_, err := s.DB.Query(query, updateAcc.FirstName, updateAcc.LastName, id)

	return err
}
func (s *PostgresStore) DeleteAccount(id int) error {
	query := "delete from account where id=$1"
	result, err := s.DB.Exec(query, id)
	if count, _ := result.RowsAffected(); count == 0 {
		return fmt.Errorf("no account found")
	}
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

func (s *PostgresStore) GetAllAccounts() ([]*Account, error) {

	accounts := []*Account{}
	var er error
	query := "select * from account"
	if rows, err := s.DB.Query(query); err != nil {
		er = err
		log.Println("get all accounts error: " + err.Error())
	} else {
		for rows.Next() {
			account := new(Account)
			rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.Created_AT)
			accounts = append(accounts, account)

		}

	}
	return accounts, er
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{FirstName: firstName, LastName: lastName, Number: rand.Intn(9999) + 1001}
}
