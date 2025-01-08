package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
	"strconv"
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

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	htmlTemplatePath = "templates/game.html"
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
	template *template.Template
	playGame Game
}

func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	p := new(PlayerServer)
	tmplt, err := template.ParseFiles(htmlTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("problem loading template: %s", err)
	}
	p.playGame = game
	p.template = tmplt
	p.Store = store
	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/game", http.HandlerFunc(p.game))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))
	p.Handler = router
	return p, nil
}

func (p *PlayerServer) game(w http.ResponseWriter, r *http.Request) {
	p.template.Execute(w, nil)
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(w, r)

	numberOfPlayersMessage := ws.WaitForMsg()
	numberOfPlayers, _ := strconv.Atoi(numberOfPlayersMessage)
	p.playGame.Start(numberOfPlayers, ws)

	winner := ws.WaitForMsg()
	p.Store.RecordWin(winner)
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
