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
	game := poker.NewTexasHoldem(poker.BlindAlertFunc(poker.Alerter), store)
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", server))
}
