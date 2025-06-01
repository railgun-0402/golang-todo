package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/models"
	"todo/services"

	"github.com/gorilla/mux"
)

var todos []models.Todo


// タスク一覧を取得する
func GetTodos(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

// タスクを登録する
func CreateTodo(w http.ResponseWriter, req *http.Request) {
	var todo models.Todo

	// Requestの中身をTodoに変換し、JSON形式で返却
	if err := json.NewDecoder(req.Body).Decode(&todo); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo);
}

// タスクをidで取得する
func GetTodoById(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	todo, err := services.GetTodoByIdService(todoID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

// タスクを更新する
func Update(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	var updatedTodo models.Todo
	if err := json.NewDecoder(req.Body).Decode(&updatedTodo); err != nil {
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


// タスクを削除する
func Delete(w http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// 「id」の要素を削除
	todos = append(todos[:id-1], todos[id:]...)
	json.NewEncoder(w).Encode(todos)
}
