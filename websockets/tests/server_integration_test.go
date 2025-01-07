package tests

import (
	"buildingHTTPServer/poker"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsThenRetrievingEm(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()
	store, err := poker.NewFSPlayerStore(database)
	assertNoError(t, err)
	server, _ := poker.NewPlayerServer(store)
	playerName := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(playerName))
	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(playerName))
		assertEqual(t, response.Code, http.StatusOK)
		assertEqual(t, response.Body.String(), "3")
	})
	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLeagueRequest())
		got := getLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}
