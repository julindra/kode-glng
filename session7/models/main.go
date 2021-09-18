package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var (
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASS")
		host   = os.Getenv("DB_HOST")
		dbname = os.Getenv("DB_NAME")
	)
	dsn := user + ":" + pass + "@tcp(" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal("Error conneting to database")
	}

	db.AutoMigrate(&Order{}, &Item{})
}

func GetDB() *gorm.DB {
	return db
}
