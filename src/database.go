package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func InitDatabase() *Database {
	db, err := sql.Open("postgres", "user=postgres host=127.0.0.1 dbname=what-to-code sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return &Database{DB: db}
}

func (db *Database) GetIdeasFromTags(tags []string) {
	//db.DB.Exec("")
}
