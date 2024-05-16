package main

func main() {
	store, _ := NewPostgresStore()
	server := NewServer(":2000", store)
	server.RunsServer()

}
