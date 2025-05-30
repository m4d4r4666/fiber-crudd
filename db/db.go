package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	connStr := "user=postgres dbname=tasks sslmode=disable password=yourpass"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return fmt.Errorf("error connecting DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("can not connect to DB: %v", err)
	}

	DB = db
	log.Println("Success connection to DB")
	return nil
}
