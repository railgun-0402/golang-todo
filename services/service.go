package services

import "database/sql"

// Service構造体
type TodoService struct {
	db *sql.DB
}

func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{db: db}
}
