package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"todo/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)



func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/get", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/get/{id:[0-9]+}", handlers.GetTodoById).Methods("GET")
	r.HandleFunc("/create", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/update/{id:[0-9]+}", handlers.Update).Methods("PUT")
	r.HandleFunc("/delete/{id:[0-9]+}", handlers.Delete).Methods("DELETE")

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
	}).Handler(r)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
