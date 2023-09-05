package main

import (
	"net/http"
	"trade_tally/internal/models"
	"trade_tally/internal/validator"
)

func (app *application) signupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.renderPage(w, http.StatusOK, "signup.tmpl", nil)
	case http.MethodPost:
		firstName := r.Form.Get("first-name")
		lastName := r.Form.Get("last-name")
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		user := &models.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		}
		err := user.Password.Set(password)
		if err != nil {
			app.serverError(w, err)
			return
		}

		v := validator.New()

		if models.ValidateUser(v, user); !v.Valid() {
			app.validationError(w, v.Errors)
			return
		}

		err = app.models.Users.Insert(user)
		if err != nil {
			app.serverError(w, err)
			return
		}
	default:
		w.Header().Set("Accept", "GET, POST")
	}
}
