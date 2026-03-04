package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// code.dev/

func main() {
	InitDatabase()
	server := Server{Port: 8080}
	router := http.NewServeMux()
	router.HandleFunc("POST /post-idea", PostIdea)
	router.HandleFunc("POST /post-review", PostReview)
	router.HandleFunc("GET /ideas", GetIdeas)

	log.Println("server running on port 8080")

	err := http.ListenAndServe(fmt.Sprintf(":%v", server.Port), router)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else {
		log.Fatal(err)
	}
}
