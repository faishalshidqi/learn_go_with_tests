package tests

import (
	"buildingHTTPServer/poker"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []poker.Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague(order bool) poker.League {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
		[]poker.Player{},
	}
	server := poker.NewPlayerServer(&store)
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"
		assertEqual(t, response.Code, http.StatusOK)
		assertEqual(t, want, got)
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"
		assertEqual(t, response.Code, http.StatusOK)
		assertEqual(t, got, want)
	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound
		assertEqual(t, got, want)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
		[]poker.Player{},
	}
	server := poker.NewPlayerServer(&store)
	t.Run("returns accepted on POST", func(t *testing.T) {
		playerName := "Pepper"
		request := newPostWinRequest(playerName)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertEqual(t, response.Code, http.StatusAccepted)
		assertEqual(t, len(store.winCalls), 1)
		assertEqual(t, store.winCalls[0], playerName)
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []poker.Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Troy", 14},
	}
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
		wantedLeague,
	}
	server := poker.NewPlayerServer(&store)

	t.Run("returns 200 on /league as json", func(t *testing.T) {
		request := newGetLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := getLeagueFromResponse(t, response.Body)

		assertEqual(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		contentType := response.Result().Header.Get("content-type")
		assertEqual(t, contentType, "application/json")
	})
}

func assertLeague(t *testing.T, got, wantedLeague []poker.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, wantedLeague) {
		t.Errorf("wanted: %v, got: %v", wantedLeague, got)
	}
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []poker.Player) {
	league, _ = poker.NewLeague(body)
	return
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertEqual[T comparable](t *testing.T, got T, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls, want %d", len(store.winCalls), 1)
	}
	if store.winCalls[0] != winner {
		t.Errorf("did not store corrent winner got %q, want %q", store.winCalls[0], winner)
	}
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}
func newGetLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}
