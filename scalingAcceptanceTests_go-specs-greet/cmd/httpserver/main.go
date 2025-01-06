package main

import (
	"go-specs-greet/adapters/httpserver"
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", httpserver.NewHandler()); err != nil {
		log.Fatal(err)
	}
}
