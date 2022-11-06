package routes

import (
	"github.com/gorilla/mux"
	"github.com/rajath-dev/todo/pkg/controllers"
)

func RegisterTodoRoutes(r *mux.Router) {
	r.HandleFunc("/todos", controllers.GetAllTodos)
}
