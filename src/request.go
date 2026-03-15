package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Response[T any] struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func Docs(w http.ResponseWriter, r *http.Request) {
	SendJSON(w, Success(http.StatusOK, "Go to \"https://github.com/MaxChip101/what-to-code\" for documentation"))
}

func PostIdea(w http.ResponseWriter, r *http.Request) {
	var idea Idea
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	err := json.NewDecoder(r.Body).Decode(&idea)
	if err != nil {
		SendJSON(w, Error(http.StatusInternalServerError, "Failed to parse idea"))
		return
	}

	if idea.Title == "" || idea.Content == "" {
		SendJSON(w, Error(http.StatusBadRequest, "No title or content provided"))
		return
	}

	err = PostIdeaIntoDB(&idea)

	if err != nil {
		SendJSON(w, Error(http.StatusInternalServerError, "Failed to post idea into database"))
		log.Println(err)
		return
	}

	SendJSON(w, Success(http.StatusOK, "success"))
}

func GetIdeas(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit_string := query.Get("limit")
	tag_string := query.Get("tags")
	id_string := query.Get("id")

	if id_string != "" {
		id, err := strconv.Atoi(id_string)
		if err != nil {
			SendJSON(w, Error(http.StatusInternalServerError, "Failed to parse id"))
			return
		}
		idea, err := GetIdeaFromId(id)
		if err != nil {
			if err == sql.ErrNoRows {
				SendJSON(w, Error(http.StatusNotFound, "No posted idea with this id"))
				return
			}

			SendJSON(w, Error(http.StatusInternalServerError, "Failed to get idea from database"))
			log.Println(err)
			return
		}
		SendJSON(w, Success(http.StatusOK, idea))
		return
	}

	var err error
	var limit int
	tags := strings.Split(tag_string, ",")
	if limit_string != "" {
		limit, err = strconv.Atoi(limit_string)
		if err != nil {
			SendJSON(w, Error(http.StatusInternalServerError, "Failed to parse limit"))
			return
		}
	} else {
		limit = 1
	}

	var ideas []Idea

	if tag_string != "" {
		ideas, err = GetIdeasFromTags(tags, limit)
		if err != nil {
			SendJSON(w, Error(http.StatusInternalServerError, "Failed to get ideas from database"))
			log.Println(err)
			return
		}
	} else {
		ideas, err = GetIdeasFromDB(limit)
		if err != nil {
			SendJSON(w, Error(http.StatusInternalServerError, "Failed to get ideas from database"))
			log.Println(err)
			return
		}
	}

	SendJSON(w, Success(http.StatusOK, ideas))
}

func SendJSON(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Error(code int, message string) Response[any] {
	return Response[any]{
		Status:  "error",
		Code:    code,
		Message: message,
	}
}

func Success[T any](code int, data T) Response[T] {
	return Response[T]{
		Status: "success",
		Code:   code,
		Data:   data,
	}
}
