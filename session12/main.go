package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"session12/db"
	TodoApi "session12/todo/api"
	TodoController "session12/todo/controller"
	TodoStorage "session12/todo/storage"

	UserApi "session12/user/api"
	UserController "session12/user/controller"
	UserStorage "session12/user/storage"

	StatusStorage "session12/status/storage"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "session12/docs"
)

// @title Todo Application
// @version 1.0
// @description Simple Todo REST API

// @contact.name Renhard Julindra

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConn := db.InitDB()

	r := gin.Default()

	todoStorage := &TodoStorage.Storage{DB: dbConn}
	userStorage := &UserStorage.Storage{DB: dbConn}
	statusStorage := &StatusStorage.Storage{DB: dbConn}

	todoController := &TodoController.Controller{Storage: todoStorage, StatusStorage: statusStorage}
	userController := &UserController.Controller{Storage: userStorage}

	docs.SwaggerInfo.BasePath = "/api/v1"
	api := r.Group("/api/v1")
	{
		TodoApi.Router(api, todoController)
		UserApi.Router(api, userController)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
