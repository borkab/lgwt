package main

import (
	"log"
	"net/http"
)

/*
// InMemoryPlayerStore collects data about players in memory
type InMemoryPlayerStore struct{}

// GetPlayerScore retrives scores for a given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}
*/

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
