package routers

import (
	"github.com/gin-gonic/gin"
)

func RouterRegister(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	TodoRouters(router)
}
