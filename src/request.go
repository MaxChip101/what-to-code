package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	Status bool   `json:"status"`
	Data   any    `json:"data,omitempty"`
	Error  string `json:"error,omitempty"`
}

func PostIdea(w http.ResponseWriter, r *http.Request) {
	var idea Idea
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	err := json.NewDecoder(r.Body).Decode(&idea)
	if err != nil {
		SendJSON(w, http.StatusInternalServerError, &Response{Status: false, Error: "failed to decode incoming json data"})
		return
	}

	err = PostIdeaIntoDB(&idea)
	if err != nil {
		SendJSON(w, http.StatusInternalServerError, &Response{Status: false, Error: "internal server error"})
		log.Println(err)
		return
	}
}

func GetIdeas(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit_string := query.Get("limit")
	tag_string := query.Get("tags")
	id_string := query.Get("id")

	// make tag searching, ignore tag searching if id provided

	if id_string != "" {
		id, err := strconv.Atoi(id_string)
		if err != nil {
			SendJSON(w, http.StatusBadRequest, &Response{Status: false, Error: "could not parse id"})
			return
		}
		idea, err := GetIdeaFromId(id)
		if err != nil {
			if err == sql.ErrNoRows {
				SendJSON(w, http.StatusNotFound, &Response{Status: false, Error: "no posted idea with this id"})
				return
			}

			SendJSON(w, http.StatusInternalServerError, &Response{Status: false, Error: "internal server error"})
			log.Println(err)
			return
		}
		SendJSON(w, http.StatusOK, &Response{Status: true, Data: []Idea{idea}})
		return
	}

	if limit_string != "" && tag_string != "" {
		tags := strings.Split(tag_string, ",")
		limit, err := strconv.Atoi(limit_string)
		if err != nil {
			SendJSON(w, http.StatusBadRequest, &Response{Status: false, Error: "could not parse limit"})
			return
		}
		ideas, err := GetIdeasFromTags(tags, limit)
		if err != nil {
			SendJSON(w, http.StatusInternalServerError, &Response{Status: false, Error: "internal server error"})
			log.Println(err)
			return
		}

		SendJSON(w, http.StatusOK, &Response{Status: true, Data: ideas})
	} else if tag_string != "" {
		tags := strings.Split(tag_string, ",")
		ideas, err := GetIdeasFromTags(tags, 1)
		if err != nil {
			SendJSON(w, http.StatusInternalServerError, &Response{Status: false, Error: "internal server error"})
			log.Println(err)
			return
		}

		SendJSON(w, http.StatusOK, &Response{Status: true, Data: ideas})
	} else {
		limit := 1
		var err error
		if limit_string != "" {
			limit, err = strconv.Atoi(limit_string)
			if err != nil {
				SendJSON(w, http.StatusBadRequest, &Response{Status: false, Error: "could not parse limit"})
				return
			}
		}

		ideas, err := GetIdeasFromDB(limit)
		if err != nil {
			SendJSON(w, http.StatusInternalServerError, &Response{Status: false, Error: "internal server error"})
			log.Println(err)
			return
		}

		SendJSON(w, http.StatusOK, &Response{Status: true, Data: ideas})
	}
}

func SendJSON(w http.ResponseWriter, status int, response *Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(response)
}
