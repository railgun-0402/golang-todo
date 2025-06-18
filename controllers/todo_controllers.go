package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/apperrors"
	"todo/controllers/services"
	"todo/models"

	"github.com/labstack/echo/v4"
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
func (c *TodoController) GetTodos(ctx echo.Context) error {
	todos, err := c.service.GetTodos()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string {
			"error": "fail internal exec",
		})
	}
	return ctx.JSON(http.StatusOK, todos)
}

// タスクを登録する
func (c *TodoController) CreateTodo(ctx echo.Context) error {
	var todo models.Todo

	// Requestの中身をTodoに変換し、JSON形式で返却
	if err := ctx.Bind(&todo); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(ctx, err)
	}

	result, err := c.service.Insert(todo)
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}
	return ctx.JSON(http.StatusOK, result)
}

// タスクをidで取得する
func (c *TodoController) GetTodoByIdHandle(ctx echo.Context) error {
	// Validate Check
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		return apperrors.ErrorHandler(ctx, err)
	}

	todo, err := c.service.GetTodoById(todoID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string {
			"error": "fail internal exec",
		})
	}

	return ctx.JSON(http.StatusOK, todo)
}

// タスクを更新する
func (c *TodoController) Update(ctx echo.Context) error {
	// Validate Check
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		apperrors.ErrorHandler(ctx, err)
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(ctx.Request().Body).Decode(&updatedTodo); err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}

	err = c.service.Update(todoID, updatedTodo.Done)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string {
			"error": "fail internal exec",
		})
	}
	return ctx.JSON(http.StatusOK, updatedTodo)
}

// タスクを削除する
func (c *TodoController) Delete(ctx echo.Context) error {

	// Validate Check
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		return apperrors.ErrorHandler(ctx, err)
	}

	err = c.service.Delete(id)
	if err != nil {
		return apperrors.ErrorHandler(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// ヘルスチェック用
func (c *TodoController) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}
