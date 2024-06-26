package main

import (
	"github.com/coregatekit/gotodoapp/common"
	"github.com/coregatekit/gotodoapp/models"
	"github.com/coregatekit/gotodoapp/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	common.LoadEnvironment()
	models.ConnectDatabase()
	models.MigrateDatabase()
}

func main() {
	r := gin.Default()

	api := r.Group("/api")
	routers.RouterRegister(api)

	r.Run(":8080")
}
