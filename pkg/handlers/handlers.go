package handlers

import (
	"html/template"
	"net/http"

	"github.com/storyteller23/todolist-go/pkg/config"
	"github.com/storyteller23/todolist-go/pkg/controller"
)

var (
	tmpl = template.Must(template.ParseFiles("ui/html/index.html"))
	db   = config.Database()
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	data, err := controller.GetTodoList(db)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	tmpl.Execute(w, data)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	r.ParseForm()

	title, ok := r.Form["title"]
	if !ok {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	controller.Add(title[0], db)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
