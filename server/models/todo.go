package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoHandler struct {
	db *gorm.DB
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db}
}

func (t *TodoHandler) GetAllTodos() ([]Todo, error) {
	var todos []Todo
	t.db.Find(&todos)

	if t.db.Error != nil {
		return nil, t.db.Error
	}

	return todos, nil
}
