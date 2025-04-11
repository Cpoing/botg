package main

// 305

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
  "time"

  "api/internal/models"
	_ "github.com/go-sql-driver/mysql"
  "github.com/alexedwards/scs/v2"
  "github.com/alexedwards/scs/mysqlstore"
	"github.com/joho/godotenv"
)

type application struct {
	logger *slog.Logger
  blogs *models.BlogModel
  users *models.UserModel
  sessionManager *scs.SessionManager
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

  sessionManager := scs.New()
  sessionManager.Store = mysqlstore.New(db)
  sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		logger: logger,
    blogs: &models.BlogModel{DB: db},
    users: &models.UserModel{DB: db},
    sessionManager: sessionManager,
	}

  srv := &http.Server{
    Addr: *addr,
    Handler: app.routes(),
    ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
    IdleTimeout: time.Minute,
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
  }

	logger.Info("Starting server", "addr", srv.Addr)

	err = srv.ListenAndServe()
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
