package webserver

import (
	"embed"
	"go-specs-greet/domain/interactions"
	"html/template"
	"net/http"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

var (
	//go:embed "markup/*"
	templates embed.FS
)

func NewHandler() (http.Handler, error) {
	tmplt, err := template.ParseFS(templates, "markup/*.gohtml")
	if err != nil {
		return nil, err
	}
	handlr := handler{tmplt}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlr.form)
	mux.HandleFunc(greetPath, handlr.replyWith(interactions.Greet))
	mux.HandleFunc(cursePath, handlr.replyWith(interactions.Curse))
	return mux, nil
}

type handler struct {
	template *template.Template
}

func (h handler) replyWith(interact func(name string) string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.template.ExecuteTemplate(w, "reply.gohtml", interact(r.Form.Get("name"))); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h handler) form(w http.ResponseWriter, _ *http.Request) {
	_ = h.template.ExecuteTemplate(w, "form.gohtml", nil)
}
