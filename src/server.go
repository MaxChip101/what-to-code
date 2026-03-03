package main

type Server struct {
	Port  int
	Ideas []Idea
}

type Idea struct {
	Creator  User      `json:"creator"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Tags     []Tag     `json:"tags"`
	Id       string    `json:"id"` // server generated
	Comments []Comment `json:"comments"`
	Requests int       `json:"requests"` // server generated, used for popularity
}

type Comment struct {
	Creator User   `json:"creator"`
	Content string `json:"content"`
	Id      string `json:"id"` // server generated
}

type Tag struct {
	Name string `json:"name"`
}

type User struct {
	Name string `json:"name"`
}

// http://localhost:8080/ideas/?id=%2212ehfwj93%22
