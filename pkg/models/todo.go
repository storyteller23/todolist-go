package models

type Todo struct {
	Id        int
	Title     string
	Completed int
}

type TodoList struct {
	Todos []Todo
}
