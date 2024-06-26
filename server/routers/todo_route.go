package routers

import (
	"github.com/coregatekit/gotodoapp/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRouters(router *gin.RouterGroup) {
	router.GET("/todos", controllers.GetAllTodos)
}
