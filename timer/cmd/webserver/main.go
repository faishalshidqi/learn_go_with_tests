package main

import (
	"buildingHTTPServer/poker"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFile, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile()
	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8080", server))
}
