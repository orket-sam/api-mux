package main

import "log"

func main() {
	server := NewAPIServer(":3000")
	log.Println("server is up and running")
	server.Runserver()

}
