package api

import (
	"database/sql"
	"todo/api/middlewares"
	"todo/controllers"
	"todo/services"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewTodoService(db)
	todoCon := controllers.NewTodoController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/get", todoCon.GetTodos).Methods("GET")
	r.HandleFunc("/get/{id:[0-9]+}", todoCon.GetTodoByIdHandle).Methods("GET")
	r.HandleFunc("/create", todoCon.CreateTodo).Methods("POST")
	r.HandleFunc("/update/{id:[0-9]+}", todoCon.Update).Methods("PUT")
	r.HandleFunc("/delete/{id:[0-9]+}", todoCon.Delete).Methods("DELETE")

	// ルータ r に登録されているハンドラの前処理・後処理として
	// LoggingMiddleware が使われるようになる
	r.Use(middlewares.LoggingMiddleware)

	return r
}