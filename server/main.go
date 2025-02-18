package main

import (
  "database/sql"
  "encoding/json"
  "log"
  "net/http"
  "os"

  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
)

func main() {
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  _, err := db.Exec("CREATE TABLE IF NOT EXISTS posts (id SERIAL PRIMARY KEY, content TEXT)")
  if err != nil {
    log.Fatal(err)
  }

  router := mux.NewRouter()
  router.HandleFunc("/api/go/posts", getUsers(db)).Mehtods("GET")

  enhancedRouter := enableCORS(jsonContentTypeMiddleware(router))

  log.Fatal(http.ListenAndServe(":8000", enhancedRouter))
}

func enableCORS(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Origin", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Origin", "Content-Type, Authorization")

    if r.Method == "OPTIONS" {
      w.WriteHeader(http.StatusOK)
      return
    }
    
    next.ServeHTTP(w, r)
  })
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
  return http.HandleFunc(func(w http.ResponseWirter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    next.ServeHTTP(w, r)
  }) 
}
