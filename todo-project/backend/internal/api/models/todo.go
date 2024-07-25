package models

type NewTodo struct {
	Title string `json:"title"`
}

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
