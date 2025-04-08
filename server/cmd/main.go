package main

// 112

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

  "api/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type application struct {
	logger *slog.Logger
  blogs *models.BlogModel
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load("../.env")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	databaseURL := os.Getenv("DATABASE_URL")

	addr := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", databaseURL, "MySQL data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger: logger,
    blogs: &models.BlogModel{DB: db},
	}

	logger.Info("Starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
