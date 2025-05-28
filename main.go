package main

import (
	"log"
	"net/http"
	"todo/handlers"

	"github.com/gorilla/mux"
)



func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/get/{id:[0-9]+}", handlers.GetTodoById).Methods("GET")
	r.HandleFunc("/create", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/update/{id:[0-9]+}", handlers.Update).Methods("PUT")
	r.HandleFunc("/delete/{id:[0-9]+}", handlers.Delete).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
