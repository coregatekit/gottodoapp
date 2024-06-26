package controllers

import (
	"net/http"

	"github.com/coregatekit/gotodoapp/models"
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
