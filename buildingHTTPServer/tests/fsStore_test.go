package tests

import (
	"buildingHTTPServer/src"
	"buildingHTTPServer/src/fileSystem"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 100}]`)
		defer cleanDatabase()
		store, err := fileSystem.NewFSPlayerStore(database)
		assertNoError(t, err)
		got := store.GetLeague(false)
		want := []src.Player{
			{"Cleo", 10},
			{"Chris", 100},
		}
		got = store.GetLeague(false)
		assertLeague(t, got, want)
	})
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 100}]`)
		defer cleanDatabase()
		store, err := fileSystem.NewFSPlayerStore(database)
		assertNoError(t, err)
		got := store.GetPlayerScore("Chris")
		want := 100
		assertEqual(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 100}]`)
		defer cleanDatabase()
		store, err := fileSystem.NewFSPlayerStore(database)
		assertNoError(t, err)
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 101
		assertEqual(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 100}]`)
		defer cleanDatabase()
		store, err := fileSystem.NewFSPlayerStore(database)
		assertNoError(t, err)
		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		assertEqual(t, got, want)
	})
	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, ``)
		defer cleanDatabase()
		_, err := fileSystem.NewFSPlayerStore(database)
		assertNoError(t, err)
	})
	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[{"Name": "Cleo", "Wins": 10},{"Name": "Chris", "Wins": 100}]`)
		defer cleanDatabase()
		store, err := fileSystem.NewFSPlayerStore(database)
		assertNoError(t, err)
		got := store.GetLeague(true)
		want := []src.Player{
			{"Chris", 100},
			{"Cleo", 10},
		}
		got = store.GetLeague(true)
		assertLeague(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("error not expected: %v", err)
	}
}
