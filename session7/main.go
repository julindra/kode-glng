package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"session7/controllers"
	"session7/middlewares"
	"session7/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	models.InitDB()

	r := gin.Default()

	r.Use(middlewares.Error())

	orders := r.Group("/orders")
	{
		orders.GET("/", controllers.Read)
		orders.POST("/", middlewares.Parser, controllers.Create)
		orders.PUT("/", middlewares.Parser, controllers.Update)
		orders.DELETE("/:id", controllers.Delete)
	}

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
