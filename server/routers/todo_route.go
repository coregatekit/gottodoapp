package routers

import (
	"github.com/coregatekit/gotodoapp/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRouters(router *gin.RouterGroup) {
	router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todos", controllers.NewTodo)
	router.PATCH("/todos/:id", controllers.UpdateTodoCompleted)
	router.DELETE("/todos/:id", controllers.DeleteTodo)
}
