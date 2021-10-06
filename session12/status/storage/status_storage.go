package storage

import (
	"session12/models"

	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func (s Storage) Read() ([]models.Status, error) {
	var statuses []models.Status
	if err := s.DB.Find(&statuses).Error; err != nil {
		return nil, err
	} else {
		return statuses, nil
	}
}

func (s Storage) ReadOneByStatus(txt string) (models.Status, error) {
	var status models.Status
	if err := s.DB.Where("status_txt = ?", txt).First(&status).Error; err != nil {
		return models.Status{}, err
	} else {
		return status, nil
	}
}
