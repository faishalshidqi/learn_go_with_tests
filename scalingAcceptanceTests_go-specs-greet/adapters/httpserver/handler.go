package httpserver

import (
	"fmt"
	"go-specs-greet/domain/interactions"
	"net/http"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(cursePath, replyWith(interactions.Curse))
	mux.HandleFunc(greetPath, replyWith(interactions.Greet))
	return mux
}

func replyWith(f func(name string) (interaction string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprintf(w, f(name))
	}
}
