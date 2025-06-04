package routers

import (
	"todo/controllers"

	"github.com/gorilla/mux"
)

func NewRouter(con *controllers.TodoController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/get", con.GetTodos).Methods("GET")
	r.HandleFunc("/get/{id:[0-9]+}", con.GetTodoByIdHandle).Methods("GET")
	r.HandleFunc("/create", con.CreateTodo).Methods("POST")
	r.HandleFunc("/update/{id:[0-9]+}", con.Update).Methods("PUT")
	r.HandleFunc("/delete/{id:[0-9]+}", con.Delete).Methods("DELETE")

	return r
}