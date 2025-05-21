package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo

// タスク一覧を取得する
func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

// タスクを登録する
func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	// Requestの中身をTodoに変換し、JSON形式で返却
	json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

// タスクをidで取得する
func getTodoById(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(todos[todoID - 1])
}

// todo: タスクを更新する
// todo: タスクを削除する

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get", getTodos).Methods("GET")
	r.HandleFunc("/get/{id:[0-9]+}", getTodoById).Methods("GET")
	r.HandleFunc("/create", createTodo).Methods("POST")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
