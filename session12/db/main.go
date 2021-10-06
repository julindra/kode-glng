package db

import (
	"log"
	"os"
	"session12/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var (
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASS")
		host   = os.Getenv("DB_HOST")
		dbname = os.Getenv("DB_NAME")
	)
	dsn := user + ":" + pass + "@tcp(" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error database connection", err)
	}

	db.AutoMigrate(&models.User{}, &models.Status{}, &models.Todo{})
	if rows := db.First(&models.Status{}).RowsAffected; rows < 1 {
		db.Create(&[]models.Status{
			{StatusTxt: "New"},
			{StatusTxt: "OnGoing"},
			{StatusTxt: "Done"},
			{StatusTxt: "Deleted"},
		})
	}
	return db
}
