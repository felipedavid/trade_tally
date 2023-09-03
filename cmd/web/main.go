package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	addr string
	port int
}

type application struct {
	config
	log *slog.Logger
}

func main() {
	var app application

	flag.StringVar(&app.addr, "addr", "127.0.0.1", "network interface the http server will run on")
	flag.IntVar(&app.port, "port", 8080, "port the http server will bind to")
	flag.Parse()

	app.log = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app.log.Info("starting web server")
	s := http.Server{
		Handler: app.routes(),
		Addr:    fmt.Sprintf("%s:%d", app.addr, app.port),
	}

	err := s.ListenAndServe()
	if err != nil {
		app.log.Error("unable to start server", "error", err)
		return
	}
}
