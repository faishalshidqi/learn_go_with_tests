package main

import (
	"buildingHTTPServer/src"
	"buildingHTTPServer/src/fileSystem"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	//store := inMemory.NewInMemoryPlayerStore()
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file %s: %v", dbFileName, err)
	}
	store, err := fileSystem.NewFSPlayerStore(db)
	if err != nil {
		log.Fatalf("error creating file system player store %s: %v", dbFileName, err)
	}
	server := src.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8080", server))
}
