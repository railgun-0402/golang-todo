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

// タスクを更新する
func update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Done = updatedTodo.Done
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}


// todo: タスクを削除する

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get", getTodos).Methods("GET")
	r.HandleFunc("/get/{id:[0-9]+}", getTodoById).Methods("GET")
	r.HandleFunc("/create", createTodo).Methods("POST")
	r.HandleFunc("/update/{id:[0-9]+}", update).Methods("PUT")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
