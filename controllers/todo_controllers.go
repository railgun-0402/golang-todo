package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/apperrors"
	"todo/controllers/services"
	"todo/models"

	"github.com/gorilla/mux"
)

// Controller構造体
type TodoController struct {
	service services.TodoAppServicer
}

// コンストラクタ
func NewTodoController(s services.TodoAppServicer) *TodoController {
	return &TodoController{service: s}
}

var todos []models.Todo

// タスク一覧を取得する
func (c *TodoController) GetTodos(w http.ResponseWriter, req *http.Request) {
	todos, err := c.service.GetTodos()
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

// タスクを登録する
func (c *TodoController) CreateTodo(w http.ResponseWriter, req *http.Request) {
	var todo models.Todo

	// Requestの中身をTodoに変換し、JSON形式で返却
	if err := json.NewDecoder(req.Body).Decode(&todo); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	result, err := c.service.Insert(todo)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(result);
}

// タスクをidで取得する
func (c *TodoController) GetTodoByIdHandle(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := c.service.GetTodoById(todoID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

// タスクを更新する
func (c *TodoController) Update(w http.ResponseWriter, req *http.Request) {
	todoID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(req.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.Update(todoID, updatedTodo.Done)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

// タスクを削除する
// todo: Service層作るの忘れた
func (c *TodoController) Delete(w http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 「id」の要素を削除
	todos = append(todos[:id-1], todos[id:]...)
	json.NewEncoder(w).Encode(todos)
}
