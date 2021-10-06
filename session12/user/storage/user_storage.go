package storage

import (
	"session12/models"

	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func (s Storage) Create(user models.User) (models.User, error) {
	if err := s.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	} else {
		return user, nil
	}
}

func (s Storage) Read() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (s Storage) ReadOne(id int) (models.User, error) {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	} else {
		return user, nil
	}
}

func (s Storage) Update(user models.User) error {
	if err := s.DB.Omit("created_at", "deleted_at").Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s Storage) Delete(id int) error {
	if err := s.DB.Unscoped().Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
