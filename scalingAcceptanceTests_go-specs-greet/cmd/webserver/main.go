package main

import (
	"go-specs-greet/adapters/webserver"
	"log"
	"net/http"
)

func main() {
	handler, err := webserver.NewHandler()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8081", handler))
}
