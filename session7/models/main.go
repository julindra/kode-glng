package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBInit() {
	dsn := "root:root@tcp(localhost:3306)/orders_by?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&Order{}, &Item{})
}

func GetDB() *gorm.DB {
	return db
}
