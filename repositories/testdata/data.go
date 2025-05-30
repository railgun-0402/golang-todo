package testdata

import "todo/models"

var SelectDetailTestData = []models.Todo {
	models.Todo {
		ID:    "1",
		Title: "firstTodo",
		Done:  false,
	},
	models.Todo{
		ID:    "2",
		Title: "secondTodo",
		Done:  true,
	},
}