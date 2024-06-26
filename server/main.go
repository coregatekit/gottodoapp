package main

import (
	"github.com/coregatekit/gotodoapp/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	routers.RouterRegister(api)

	r.Run(":8080")
}
