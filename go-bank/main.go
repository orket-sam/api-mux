package main

func main() {
	store, _ := NewPostgresStore()
	s := NewServer(":2000", store)
	s.RunsServer()
}
