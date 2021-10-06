package storage

import (
	"session12/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Storage struct {
	DB *gorm.DB
}

func (s Storage) Create(todo models.Todo) (models.Todo, error) {
	if err := s.DB.Create(&todo).Error; err != nil {
		return models.Todo{}, err
	} else {
		return todo, nil
	}
}

func (s Storage) Read() ([]models.Todo, error) {
	var todos []models.Todo
	if err := s.DB.Preload(clause.Associations).Find(&todos).Error; err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}

func (s Storage) ReadOne(id int) (models.Todo, error) {
	var todo models.Todo
	if err := s.DB.Unscoped().Preload(clause.Associations).First(&todo, id).Error; err != nil {
		return models.Todo{}, err
	} else {
		return todo, nil
	}
}

func (s Storage) Update(todo models.Todo) error {
	if err := s.DB.Omit("created_at", "deleted_at").Save(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (s Storage) Delete(id int, statusId int) error {
	if err := s.DB.Delete(&models.Todo{}, id).Error; err != nil {
		return err
	}
	if err := s.DB.Unscoped().Model(&models.Todo{}).Where("ID = ?", id).Update("status", statusId).Error; err != nil {
		return err
	}
	return nil
}
