package main

import "log"

func main() {
	postgresDB, err := ConnectToDb()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("DB init")
	server := NewAPIServer(":3000", postgresDB)
	log.Println("Api is up and running!!")

	server.RunServer()

}
