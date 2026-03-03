package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// code.dev/

func main() {
	server := Server{Port: 8080}
	router := mux.NewRouter()
	router.HandleFunc("/", DisplayWebsite)
	router.HandleFunc("/submit-idea", SubmitIdea).Methods("POST")
	router.HandleFunc("/submit-comment", SubmitComment).Methods("POST")
	router.HandleFunc("/get-idea", GetIdea).Methods("GET")
	router.HandleFunc("/get-random-idea", GetRandomIdea).Methods("GET")
	router.HandleFunc("/get-popular-ideas", GetPopularIdea).Methods("GET")
	router.HandleFunc("/get-new-ideas", GetNewIdea).Methods("GET")
	router.HandleFunc("/get-random-idea", GetRandomIdea).Methods("GET")
	router.HandleFunc("/get-idea", GetIdea).Methods("GET")
	router.HandleFunc("/get-comment", GetComment).Methods("GET")
	router.HandleFunc("/get-comments", GetComments).Methods("GET")

	router.PathPrefix("/res/").Handler(http.StripPrefix("/res/", http.FileServer(http.Dir("./res"))))
	err := http.ListenAndServe(fmt.Sprintf(":%v", server.Port), router)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else {
		log.Fatal(err)
	}
}
