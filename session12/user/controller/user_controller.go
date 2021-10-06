package controller

import (
	"session12/models"
)

type Controller struct {
	Storage models.UserStorage
}

func (c Controller) Create(user models.User) (models.User, error) {
	newUser, err := c.Storage.Create(user)
	return newUser, err
}

func (c Controller) Read() ([]models.User, error) {
	users, err := c.Storage.Read()
	return users, err
}

func (c Controller) Update(id int, user models.User) error {
	if _, err := c.Storage.ReadOne(id); err != nil {
		return err
	} else {
		user.UserID = uint(id)
		err := c.Storage.Update(user)
		return err
	}
}

func (c Controller) Delete(id int) error {
	if _, err := c.Storage.ReadOne(id); err != nil {
		return err
	} else {
		err := c.Storage.Delete(id)
		return err
	}
}
