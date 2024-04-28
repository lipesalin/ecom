package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDBConnection(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)

	// verifica erro
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
