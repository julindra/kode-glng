package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    uint   `gorm:"primaryKey"`
	Name      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserController interface {
	Create(user User) (User, error)
	Read() ([]User, error)
	Update(id int, user User) error
	Delete(id int) error
}

type UserStorage interface {
	Create(user User) (User, error)
	Read() ([]User, error)
	ReadOne(id int) (User, error)
	Update(user User) error
	Delete(id int) error
}
