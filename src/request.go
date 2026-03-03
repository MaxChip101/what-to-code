package main

import "net/http"

type SubmitIdeaPayload struct {
}

func DisplayWebsite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("real"))
}

func SubmitIdea(w http.ResponseWriter, r *http.Request) {

}

func SubmitComment(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	println(values)
}

func GetIdea(w http.ResponseWriter, r *http.Request) {

}

func GetRandomIdea(w http.ResponseWriter, r *http.Request) {

}

func GetPopularIdea(w http.ResponseWriter, r *http.Request) {

}

func GetNewIdea(w http.ResponseWriter, r *http.Request) {

}

func GetComment(w http.ResponseWriter, r *http.Request) {

}

func GetComments(w http.ResponseWriter, r *http.Request) {

}
