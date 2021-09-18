package controllers

import (
	"session7/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Read(c *gin.Context) {
	db := models.GetDB()
	var orders []models.Order
	if err := db.Preload("Items").Find(&orders).Error; err != nil {
		panic(err.Error())
	} else {
		c.JSON(200, orders)
	}
}

func Create(c *gin.Context) {
	db := models.GetDB()
	order := c.MustGet("order").(models.Order)
	if err := db.Create(&order).Error; err != nil {
		panic(err.Error())
	} else {
		c.JSON(201, order)
	}
}

func Update(c *gin.Context) {
	db := models.GetDB()
	order := c.MustGet("order").(models.Order)
	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&order).Error; err != nil {
		panic(err.Error())
	} else {
		c.JSON(200, order)
	}
}

func Delete(c *gin.Context) {
	db := models.GetDB()
	if err := db.Delete(&models.Order{}, c.Param("id")).Error; err != nil {
		panic(err.Error())
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}
