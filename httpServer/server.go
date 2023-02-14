package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore stores information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer is an HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store
	// to fill in the http.Handler we assign it to the router we create in
	// NewPlayerServer. We can do this because http.ServeMux has the method
	// ServeHTTP.
	// This let us remove our own ServeHTTP method, as we are exposing one
	// via the embedded new interfaces.
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	return p
}
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	leagueTable := []Player{
		{"Chris", 20},
	}

	json.NewEncoder(w).Encode(leagueTable)

	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(404)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {

	p.store.RecordWin(player)
	w.WriteHeader(202)
}
