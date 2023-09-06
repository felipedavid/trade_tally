package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"trade_tally/internal/models"

	"github.com/golang-migrate/migrate/v4"

	msqlite "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

type config struct {
	addr   string
	port   int
	dbPath string
}

type application struct {
	config
	log    *slog.Logger
	models *models.Models
}

func main() {
	var app application

	flag.StringVar(&app.addr, "addr", "127.0.0.1", "network interface the http server will run on")
	flag.StringVar(&app.dbPath, "db-path", "trade_tally.db", "database file path")
	flag.IntVar(&app.port, "port", 8080, "port the http server will bind to")
	flag.Parse()

	app.log = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := openDB(app.dbPath)
	if err != nil {
		app.log.Error("unable to connect to the database", "error", err)
		return
	}

	app.models, err = models.New(db)
	if err != nil {
		app.log.Error("unable to setup models", "error", err)
		return
	}

	// Try to run migrations
	driver, err := msqlite.WithInstance(db, &msqlite.Config{})
	if err != nil {
		app.log.Error("Unable to run migrations", "error", err)
		os.Exit(-1)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "sqlite://"+app.dbPath, driver)
	if err != nil {
		app.log.Error("Unable to run migrations", "error", err)
		os.Exit(-1)
	}

	// If they are already applied, just keep going
	app.log.Info("Running migrations")
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		app.log.Error("Unable to run migrations", "error", err)
		os.Exit(-1)
	}

	app.log.Info("starting web server")
	s := http.Server{
		Handler: app.routes(),
		Addr:    fmt.Sprintf("%s:%d", app.addr, app.port),
	}

	err = s.ListenAndServe()
	if err != nil {
		app.log.Error("unable to start server", "error", err)
		return
	}
}

func openDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
