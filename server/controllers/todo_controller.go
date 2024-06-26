package controllers

import (
	"fmt"
	"net/http"

	"github.com/coregatekit/gotodoapp/models"
	"github.com/coregatekit/gotodoapp/models/validators"
	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
	td := models.NewTodoHandler(models.DB)

	todos, err := td.GetAllTodos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func NewTodo(c *gin.Context) {
	var newTodo validators.NewTodoBody
	var err error

	if err = c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo models.Todo
	td := models.NewTodoHandler(models.DB)
	todo, err = td.NewTodo(models.Todo{Title: newTodo.Title, Completed: false})

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}
