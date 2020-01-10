package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Conn() *sql.DB {
	// connstr := "user=postgres dbname=notesite host=localhost password=anteneh23 sslmode=disable"
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB_NAME")
	connstr := fmt.Sprintf("host=%s user=%s password=%S databaseName=%s sslmode=disable", host, user, password, databaseName)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()
	fmt.Println("connected")
	return db
}
