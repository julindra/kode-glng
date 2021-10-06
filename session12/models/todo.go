package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title          string    `validate:"required"`
	Description    string    `validate:"required"`
	DueDate        time.Time `validate:"required"`
	PersonInCharge uint      `validate:"required"`
	Status         uint      `validate:"required"`
	UserData       User      `gorm:"foreignKey:PersonInCharge;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StatusData     Status    `gorm:"foreignKey:Status;references:StatusID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TodoController interface {
	Create(todo Todo) (Todo, error)
	Read() ([]Todo, error)
	ReadOne(id int) (Todo, error)
	Update(id int, todo Todo) error
	Delete(id int) error
}

type TodoStorage interface {
	Create(todo Todo) (Todo, error)
	Read() ([]Todo, error)
	ReadOne(id int) (Todo, error)
	Update(todo Todo) error
	Delete(id int, statusId int) error
}
