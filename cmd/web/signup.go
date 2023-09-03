package main

import (
	"io"
	"net/http"
)

func (app *application) signupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.renderPage(w, http.StatusOK, "signup.tmpl", nil)
	case http.MethodPost:
		_, _ = io.WriteString(w, "Creating a user account...")
	default:
		w.Header().Set("Accept", "GET, POST")
	}
}
