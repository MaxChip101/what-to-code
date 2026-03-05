package main

type Server struct {
	Port int
}

type Idea struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Id      int      `json:"id"` // server generated
}

// http://localhost:8080/idea/?id=1235464
