package api

import (
	"encoding/json"
	"io/ioutil"
	"session12/helpers"
	"session12/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TodoApi struct {
	controller models.TodoController
}

func Router(r *gin.RouterGroup, controller models.TodoController) {
	handler := &TodoApi{controller}

	todos := r.Group("/todos")
	{
		todos.POST("/", handler.Create)
		todos.GET("/", handler.Read)
		todos.GET("/:id", handler.ReadOne)
		todos.PUT("/:id", handler.Update)
		todos.DELETE("/:id", handler.Delete)
	}
}

func (t TodoApi) Validate(todo models.Todo) error {
	validate := validator.New()
	return validate.StructExcept(todo, "UserData")
}

// @Summary Create a Todo
// @Accept json
// @Produce json
// @Param json body models.Todo true "Todo"
// @Success 201 {object} models.Todo
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /todos [post]
func (t TodoApi) Create(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	var todo models.Todo
	json.Unmarshal(jsonData, &todo)

	if err := t.Validate(todo); err != nil {
		helpers.Error(c, 400, err.Error())
	} else {
		todo, err := t.controller.Create(todo)
		if err != nil {
			helpers.Error(c, 500, err.Error())
		} else {
			c.JSON(201, todo)
		}
	}
}

// @Summary Get all Todos
// @Accept json
// @Produce json
// @Success 200 {array} models.Todo
// @Failure 500 {object} models.Error
// @Router /todos [get]
func (t TodoApi) Read(c *gin.Context) {
	todos, err := t.controller.Read()
	if err != nil {
		helpers.Error(c, 500, err.Error())
	} else {
		c.JSON(200, todos)
	}
}

// @Summary Get a Todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /todos/:id [get]
func (t TodoApi) ReadOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.Error(c, 400, "Params Error: id (integer)")
	} else {
		todo, err := t.controller.ReadOne(id)
		if err != nil {
			if err.Error() == "record not found" {
				helpers.Error(c, 404, "Todo Not Found")
			} else {
				helpers.Error(c, 500, err.Error())
			}
		} else {
			c.JSON(200, todo)
		}
	}
}

// @Summary Update a Todo
// @Accept json
// @Produce json
// @Param json body models.Todo true "Todo"
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /todos/:id [put]
func (t TodoApi) Update(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	var todo models.Todo
	json.Unmarshal(jsonData, &todo)

	if err := t.Validate(todo); err != nil {
		helpers.Error(c, 400, err.Error())
	} else {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			helpers.Error(c, 400, "Params Error: id (integer)")
		} else {
			if err := t.controller.Update(id, todo); err != nil {
				if err.Error() == "record not found" {
					helpers.Error(c, 404, "Todo Not Found")
				} else {
					helpers.Error(c, 500, err.Error())
				}
			} else {
				c.JSON(200, gin.H{"Message": "Success"})
			}
		}
	}
}

// @Summary Delete a Todo
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} object{Message=string}
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /todos/:id [delete]
func (t TodoApi) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.Error(c, 400, "Params Error: id (integer)")
	} else {
		if err := t.controller.Delete(id); err != nil {
			if err.Error() == "record not found" {
				helpers.Error(c, 404, "Todo Not Found")
			} else {
				helpers.Error(c, 500, err.Error())
			}
		} else {
			c.JSON(200, gin.H{"Message": "Success"})
		}
	}
}
