package main

import (
	"io"
	"net/http"
)

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.renderPage(w, http.StatusOK, "login.tmpl", nil)
	case http.MethodPost:
		_, _ = io.WriteString(w, "Authenticating user...")
	default:
		w.Header().Set("Accept", "GET, POST")
	}
}
