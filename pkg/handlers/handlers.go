package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/storyteller23/todolist-go/pkg/config"
	"github.com/storyteller23/todolist-go/pkg/controller"
	"github.com/storyteller23/todolist-go/pkg/models"
)

var (
	index      = template.Must(template.ParseFiles("ui/html/index.html"))
	updatePage = template.Must(template.ParseFiles("ui/html/update.html"))
	db         = config.Database()
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := controller.CreateTable(db)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	uncompleted, err := controller.GetUncompletedTasks(db)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	completed, err := controller.GetCompletedTasks(db)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	data := models.TodoList{
		CompletedTasks:   completed,
		UncompletedTasks: uncompleted,
	}

	index.Execute(w, data)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()

	title, ok := r.Form["title"]
	if !ok {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	controller.Add(title[0], db)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	params := mux.Vars(r)
	id := params["id"]
	if r.Method == http.MethodPost {
		r.ParseForm()
		newTitle, ok := r.Form["newTitle"]
		if !ok {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		err := controller.Update(id, newTitle[0], db)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

	updatePage.Execute(w, id)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	params := mux.Vars(r)
	id := params["id"]
	controller.Delete(id, db)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	params := mux.Vars(r)
	id := params["id"]
	controller.CompleteTask(id, db)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
