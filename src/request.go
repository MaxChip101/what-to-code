package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func PostIdea(w http.ResponseWriter, r *http.Request) {

}

func PostReview(w http.ResponseWriter, r *http.Request) {

}

func GetIdeas(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit_string := query.Get("limit")
	tag_string := query.Get("tags")
	tags := strings.Split(tag_string, ",")
	limit, err := strconv.Atoi(limit_string)
	if err != nil {
		fmt.Fprintf(w, "error parsing limit: %a", err)
		return
	}

	fmt.Fprintf(w, "tags: %v; limit: %v", tags, limit)

}
