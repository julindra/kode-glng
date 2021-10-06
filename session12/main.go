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
)

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

	api := r.Group("/api/v1")
	{
		TodoApi.Router(api, todoController)
		UserApi.Router(api, userController)
	}

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
