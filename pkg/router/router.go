package router

import (
	"github.com/gorilla/mux"
	"github.com/storyteller23/todolist-go/pkg/handlers"
)

func Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomePage)
	router.HandleFunc("/add", handlers.AddTodo)
	return router
}
