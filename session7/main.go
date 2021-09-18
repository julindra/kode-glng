package main

import (
	"github.com/gin-gonic/gin"

	"session7/controllers"
	"session7/middlewares"
	"session7/models"
)

func main() {
	models.DBInit()

	r := gin.Default()

	r.Use(middlewares.Error())

	orders := r.Group("/orders")
	{
		orders.GET("/", controllers.Read)
		orders.POST("/", middlewares.Parser, controllers.Create)
		orders.PUT("/", middlewares.Parser, controllers.Update)
		orders.DELETE("/:id", controllers.Delete)
	}

	r.Run()
}
