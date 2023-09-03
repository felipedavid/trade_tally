package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var functions = template.FuncMap{}

type templateData struct {
}

func (app *application) renderPageBlock(w http.ResponseWriter, statusCode int, page, block string, data *templateData) {
	ts, err := template.New(page).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	ts, err = ts.ParseFiles(fmt.Sprintf("./ui/html/pages/%s", page))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//ts, err = ts.ParseGlob("./ui/html/components/*.tmpl")
	//if err != nil {
	//	return
	//}

	buf := new(bytes.Buffer)

	err = ts.ExecuteTemplate(buf, block, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	_, _ = buf.WriteTo(w)
	return
}

func (app *application) renderPage(w http.ResponseWriter, statusCode int, fileName string, data *templateData) {
	app.renderPageBlock(w, statusCode, fileName, "base", data)
}
