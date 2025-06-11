package services

import "todo/models"

// Service層メソッドの型を持つインターフェース
// ControllerがServiceに依存しないよう、インターフェースを用意し
// 同じ型を持つのであれば受付可能
type TodoAppServicer interface {
	GetTodos() ([]models.Todo, error)
	GetTodoById(id int) (models.Todo, error)
	Insert(todo models.Todo) (models.Todo, error)
	Update(id int, done bool) (error)
	Delete(id int) (error)
}