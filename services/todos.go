package services

import (
	"todo/models"
	"todo/repositories"
)

// タスクを全て取得する
func (s *TodoService) GetTodos() ([]models.Todo, error) {
	todos, err := repositories.SelectTodos(s.db)
	if err != nil {
		return []models.Todo{}, err
	}

	return todos, nil
}


// タスクをidで取得する
func (s *TodoService) GetTodoById(id int) (models.Todo, error) {
	// DBからIDに紐づくタスクを取得
	todo, err := repositories.SelectDetailTodo(s.db, id)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

// タスクを追加するService関数
func (s *TodoService) Insert(todo models.Todo) (models.Todo, error) {
	result, err := repositories.InsertTodo(s.db, todo)
	if err != nil {
		return models.Todo{}, err
	}
	return result, nil
}

// タスクを更新するService関数
func (s *TodoService) Update(id int, done bool) (error) {
	err := repositories.UpdateTodo(s.db, id, done)
	if err != nil {
		return err
	}
	return nil
}