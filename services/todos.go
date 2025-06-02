package services

import (
	"todo/models"
	"todo/repositories"
)

// タスクを全て取得する
func GetTodosService() ([]models.Todo, error) {
	db, err := connectDB()
	if err != nil {
		return []models.Todo{}, err
	}
	defer db.Close()

	todos, err := repositories.SelectTodos(db)
	if err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}


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

// タスクを追加するService関数
func InsertService(todo models.Todo) (models.Todo, error) {
	db, err := connectDB()
	if err != nil {
		return models.Todo{}, err
	}
	defer db.Close()

	result, err := repositories.InsertTodo(db, todo)
	if err != nil {
		return models.Todo{}, err
	}
	return result, nil
}

// タスクを更新するService関数
func UpdateService(id int, done bool) (error) {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = repositories.UpdateTodo(db, id, done)
	if err != nil {
		return err
	}
	return nil
}