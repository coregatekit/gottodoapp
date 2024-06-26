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

func (t *TodoHandler) NewTodo(todo Todo) (Todo, error) {
	t.db.Create(&todo)

	if t.db.Error != nil {
		return Todo{}, t.db.Error
	}

	return todo, nil
}

func (t *TodoHandler) UpdateCompleted(id int) error {
	todo := Todo{}
	t.db.First(&todo, id)

	if t.db.Error != nil {
		return t.db.Error
	}

	todo.Completed = !todo.Completed
	t.db.Save(&todo)

	if t.db.Error != nil {
		return t.db.Error
	}

	return nil
}

func (t *TodoHandler) DeleteTodo(id int) error {
	todo := Todo{}
	t.db.First(&todo, id)

	if t.db.Error != nil {
		return t.db.Error
	}

	t.db.Delete(&todo)

	if t.db.Error != nil {
		return t.db.Error
	}

	return nil
}
