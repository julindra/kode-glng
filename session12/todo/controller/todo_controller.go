package controller

import "session12/models"

type Controller struct {
	Storage       models.TodoStorage
	StatusStorage models.StatusStorage
}

func (c Controller) Create(todo models.Todo) (models.Todo, error) {
	newTodo, err := c.Storage.Create(todo)
	if err == nil {
		return c.ReadOne(int(newTodo.ID))
	}
	return newTodo, err
}

func (c Controller) Read() ([]models.Todo, error) {
	todos, err := c.Storage.Read()
	return todos, err
}

func (c Controller) ReadOne(id int) (models.Todo, error) {
	todo, err := c.Storage.ReadOne(id)
	return todo, err
}

func (c Controller) Update(id int, todo models.Todo) error {
	if _, err := c.ReadOne(id); err != nil {
		return err
	} else {
		todo.ID = uint(id)
		err := c.Storage.Update(todo)
		return err
	}
}

func (c Controller) Delete(id int) error {
	if _, err := c.ReadOne(id); err != nil {
		return err
	} else {
		status, err := c.StatusStorage.ReadOneByStatus("Deleted")
		if err != nil {
			return err
		}
		err = c.Storage.Delete(id, int(status.StatusID))
		return err
	}
}
