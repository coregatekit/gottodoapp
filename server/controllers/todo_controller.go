package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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

func UpdateTodoCompleted(c *gin.Context) {
	var err error
	var id int
	idStr := c.Param("id")

	id, err = strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	td := models.NewTodoHandler(models.DB)
	err = td.UpdateCompleted(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated"})
}
