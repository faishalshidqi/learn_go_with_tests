package httpserver

import (
	"fmt"
	"go-specs-greet/domain/interactions"
	"net/http"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, interactions.Greet(name))
}

func CurseHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, interactions.Curse(name))
}
