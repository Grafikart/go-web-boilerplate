package server

import (
	"grafikart/boilerplate/templates"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	component := templates.Layout("John")
	component.Render(r.Context(), w)
}
