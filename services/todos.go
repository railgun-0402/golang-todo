package services

import (
	"todo/models"
	"todo/repositories"
)

// タスクをidで取得する
func GetTodoByIdService(id int) (models.Todo, error) {
	db, err := connectDB()
	if err != nil {
		return models.Todo{}, err
	}
	defer db.Close()

	// DBからIDに紐づくタスクを取得
	todo, err := repositories.SelectDetailTodo(db, id)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}