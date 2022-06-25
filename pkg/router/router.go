package router

import (
	"github.com/gorilla/mux"
	"github.com/storyteller23/todolist-go/pkg/handlers"
)

func Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomePage)
	router.HandleFunc("/add", handlers.AddTodo)
	router.HandleFunc("/delete/{id:[0-9]+}", handlers.DeleteTodo)
	router.HandleFunc("/complete/{id:[0-9]+}", handlers.CompleteTodo)
	router.HandleFunc("/update/{id:[0-9]+}", handlers.UpdateTodo)
	router.HandleFunc("/completed", handlers.CompletedTasks)
	return router
}
