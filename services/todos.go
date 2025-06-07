package services

import (
	"database/sql"
	"errors"
	"todo/apperrors"
	"todo/models"
	"todo/repositories"
)

// タスクを全て取得する
func (s *TodoService) GetTodos() ([]models.Todo, error) {
	todos, err := repositories.SelectTodos(s.db)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return []models.Todo{}, err
	}

	if len(todos) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return todos, nil
}


// タスクをidで取得する
func (s *TodoService) GetTodoById(id int) (models.Todo, error) {
	// DBからIDに紐づくタスクを取得
	todo, err := repositories.SelectDetailTodo(s.db, id)
	if err != nil {
		// SELECTで取得失敗とデータ0件のエラーを分ける
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Todo{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Todo{}, err
	}

	return todo, nil
}

// タスクを追加するService関数
func (s *TodoService) Insert(todo models.Todo) (models.Todo, error) {
	result, err := repositories.InsertTodo(s.db, todo)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Todo{}, err
	}
	return result, nil
}

// タスクを更新するService関数
func (s *TodoService) Update(id int, done bool) (error) {
	err := repositories.UpdateTodo(s.db, id, done)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target todo")
			return err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update todo")
		return err
	}
	return nil
}