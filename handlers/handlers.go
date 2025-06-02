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
func GetTodosHandle(w http.ResponseWriter, req *http.Request) {
	todos, err := services.GetTodosService()
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

// タスクを登録する
func CreateTodoHandle(w http.ResponseWriter, req *http.Request) {
	var todo models.Todo

	// Requestの中身をTodoに変換し、JSON形式で返却
	if err := json.NewDecoder(req.Body).Decode(&todo); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	result, err := services.InsertService(todo)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result);
}

// タスクをidで取得する
func GetTodoByIdHandle(w http.ResponseWriter, r *http.Request) {
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
func UpdateHandle(w http.ResponseWriter, req *http.Request) {
	todoID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(req.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	err = services.UpdateService(todoID, updatedTodo.Done)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

// タスクを削除する
// todo: Service層作るの忘れた
func DeleteHandle(w http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// 「id」の要素を削除
	todos = append(todos[:id-1], todos[id:]...)
	json.NewEncoder(w).Encode(todos)
}
