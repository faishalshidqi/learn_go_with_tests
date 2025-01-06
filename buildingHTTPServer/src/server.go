package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*
type InMemoryPlayerStore struct{}

	func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
		return 123
	}

func (i *InMemoryPlayerStore) RecordWin(name string) {

}
*/
const (
	jsonContentType = "application/json"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(playerName string) int
	RecordWin(playerName string)
	GetLeague(order bool) League
}

type PlayerServer struct {
	Store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.Store = store
	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			p.playersHandler(w, r)
		}))
	router.Handle("/league", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			p.leagueHandler(w, r)
		}))
	p.Handler = router
	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	leagueTable := p.Store.GetLeague(false)
	marshal, _ := json.Marshal(leagueTable)
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
	return
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, playerName)
	case http.MethodPost:
		p.processWin(w, playerName)
	}
	return
}

func (p *PlayerServer) processWin(w http.ResponseWriter, playerName string) {
	p.Store.RecordWin(playerName)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, playerName string) {
	score := p.Store.GetPlayerScore(playerName)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}
