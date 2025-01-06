package tests

import (
	"buildingHTTPServer/src"
	"buildingHTTPServer/src/fileSystem"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsThenRetrievingEm(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()
	store, err := fileSystem.NewFSPlayerStore(database)
	assertNoError(t, err)
	server := src.NewPlayerServer(store)
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
		want := []src.Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}
