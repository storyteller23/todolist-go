package models

type Todo struct {
	Id        int
	Title     string
	Completed int
}

type TodoList struct {
	UncompletedTasks []Todo
	CompletedTasks   []Todo
}
