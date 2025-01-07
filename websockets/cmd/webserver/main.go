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
	server, err := poker.NewPlayerServer(store)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}
