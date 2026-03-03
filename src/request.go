package main

import (
	"net/http"
)

func PostIdea(w http.ResponseWriter, r *http.Request) {

}

func PostComment(w http.ResponseWriter, r *http.Request) {

}

func GetIdea(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tags := query.Get("tags")
}

func GetIdeas(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit := query.Get("limit")
	tags := query.Get("tags")

}
