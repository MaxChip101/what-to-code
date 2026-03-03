package main

type Server struct {
	Port int
}

type Idea struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Id      string   `json:"id"` // server generated
	Reviews []Review `json:"reviews"`
}

type Review struct {
	Rating  bool   `json:"rating"`
	Content string `json:"content"`
	Id      string `json:"id"` // server generated
}

type Tag struct {
	Name string `json:"name"`
}

// http://localhost:8080/idea/?id=1235464
