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
	// Validate Check
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
	// Validate Check
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
func (c *TodoController) Delete(w http.ResponseWriter, req *http.Request) {

	// Validate Check
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
