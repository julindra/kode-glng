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

	r.GET("/orders", controllers.Read())
	r.POST("/orders", middlewares.Parser(), controllers.Create())
	r.PUT("/orders", middlewares.Parser(), controllers.Update())
	r.DELETE("/orders/:id", controllers.Delete())

	r.Run()
}
