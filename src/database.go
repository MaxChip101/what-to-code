package main

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() {
	var err error
	db, err = sql.Open("postgres", "user=postgres host=127.0.0.1 dbname=ideas sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func GetIdeasFromTags(tags []string, limit int) ([]Idea, error) {
	rows, err := db.Query("SELECT * FROM ideas WHERE tags @> $1 ORDER BY RANDOM() LIMIT $2", pq.Array(tags), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ideas []Idea
	for rows.Next() {
		var idea Idea
		err := rows.Scan(&idea.Id, &idea.Title, &idea.Content, pq.Array(&idea.Tags))
		if err != nil {
			return nil, err
		}
		ideas = append(ideas, idea)
	}
	return ideas, nil
}

func GetIdeasFromDB(limit int) ([]Idea, error) {
	rows, err := db.Query("SELECT * FROM ideas ORDER BY RANDOM() LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ideas []Idea
	for rows.Next() {
		var idea Idea
		err := rows.Scan(&idea.Id, &idea.Title, &idea.Content, pq.Array(&idea.Tags))
		if err != nil {
			return nil, err
		}
		ideas = append(ideas, idea)
	}
	return ideas, nil
}

func GetIdeaFromId(id int) (Idea, error) {
	var idea Idea
	err := db.QueryRow("SELECT * FROM ideas WHERE id = $1", id).Scan(&idea.Id, &idea.Title, &idea.Content, pq.Array(&idea.Tags))
	if err != nil {
		return Idea{}, err
	}
	return idea, nil
}

func PostIdeaIntoDB(idea *Idea) error {
	_, err := db.Exec("INSERT INTO ideas (title, content, tags) VALUES ($1, $2, $3)", idea.Title, idea.Content, idea.Tags)
	if err != nil {
		return err
	}
	return nil
}
