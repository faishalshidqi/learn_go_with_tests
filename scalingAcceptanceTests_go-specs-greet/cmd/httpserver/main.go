package main

import (
	"go-specs-greet/adapters/httpserver"
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}
	go serveMux.HandleFunc(
		"/greet",
		httpserver.GreetHandler,
	)
	go serveMux.HandleFunc(
		"/curse",
		httpserver.CurseHandler,
	)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
