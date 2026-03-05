package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func PostIdea(w http.ResponseWriter, r *http.Request) {

}

func GetIdeas(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit_string := query.Get("limit")
	tag_string := query.Get("tags")
	id_string := query.Get("id")
	tags := strings.Split(tag_string, ",")
	limit, err := strconv.Atoi(limit_string)
	_ = limit
	_ = tags
	if err != nil {
		SendStatusMessage(w, false, "error parsing limit: "+err.Error())
		return
	}
	if id_string != "" {
		id, err := strconv.Atoi(id_string)
		if err != nil {
			SendStatusMessage(w, false, "error parsing id: "+err.Error())
			return
		}
		idea, err := GetIdeaFromId(id)
		if err != nil {
			if err == sql.ErrNoRows {
				SendStatusMessage(w, false, "no posted idea with this id")
				return
			}
			SendStatusMessage(w, false, err.Error())
			return
		}
		data, err := json.Marshal(idea)
		if err != nil {
			SendStatusMessage(w, false, err.Error())
			return
		}
		w.Write(data)
	}

}

func SendStatusMessage(w http.ResponseWriter, status bool, message string) {
	fmt.Fprintf(w, "{\"status\":%v,\"message\":%v}", status, message)
}
