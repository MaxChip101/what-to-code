package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

var limiter *rate.Limiter

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && !limiter.Allow() {
			SendJSON(w, Error(http.StatusTooManyRequests, "Rate limit exceeded"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	InitDatabase()
	limiter = rate.NewLimiter(6, 24)
	server := Server{Port: 8080}
	router := http.NewServeMux()
	router.HandleFunc("/", Docs)
	router.HandleFunc("POST /post-idea", PostIdea)
	router.HandleFunc("GET /ideas", GetIdeas)

	log.Println("server running on port 8080")

	err := http.ListenAndServe(fmt.Sprintf(":%v", server.Port), router)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else {
		log.Fatal(err)
	}
}
