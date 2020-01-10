package database

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  "log"
)

func Conn() *sql.DB {
  connstr := "user=postgres dbname=notesite host=localhost password=anteneh23 sslmode=disable"
  db, err := sql.Open("postgres", connstr)
  if err != nil {
    log.Fatal(err)
    return nil
  }
  defer db.Close()
  fmt.Println("connected")
  return db
}
