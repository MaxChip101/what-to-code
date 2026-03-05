package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() {
	var err error
	db, err = sql.Open("postgres", "user=postgres host=127.0.0.1 dbname=what_to_code sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func GetIdeasFromTags(tags []string, limit int) {
	//db.DB.Exec("")
}

func GetIdeaFromId(id int) (*Idea, error) {
	var idea Idea
	err := db.QueryRow("SELECT id, title, content, tags FROM ideas FROM ideas WHERE id = $1", id).Scan(&idea.Id, &idea.Title, &idea.Content, &idea.Tags)
	if err != nil {
		return nil, err
	}
	return &idea, nil
}

func PostIdeaIntoDB(idea *Idea) error {
	return nil
}
