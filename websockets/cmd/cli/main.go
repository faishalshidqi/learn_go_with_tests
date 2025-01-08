package main

import (
	"buildingHTTPServer/poker"
	"fmt"
	"log"
	"os"
)

const (
	dbFileName = "game.db.json"
)

func main() {
	store, closeFile, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem creating store %v", err)
	}
	defer closeFile()
	fmt.Println("Wanna play poker?")
	fmt.Println("Type {name} wins to record a win")
	game := poker.NewTexasHoldem(poker.BlindAlertFunc(poker.Alerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
