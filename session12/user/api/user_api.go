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

type UserApi struct {
	controller models.UserController
}

func Router(r *gin.RouterGroup, controller models.UserController) {
	handler := &UserApi{controller}

	users := r.Group("/users")
	{
		users.POST("/", handler.Create)
		users.GET("/", handler.Read)
		users.PUT("/:id", handler.Update)
		users.DELETE("/:id", handler.Delete)
	}
}

func (u UserApi) Validate(user models.User) error {
	validate := validator.New()
	return validate.Struct(user)
}

func (u UserApi) Create(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	var user models.User
	if err := json.Unmarshal(jsonData, &user); err != nil {
		helpers.Error(c, 400, err.Error())
	} else if err := u.Validate(user); err != nil {
		helpers.Error(c, 400, err.Error())
	} else {
		user, err := u.controller.Create(user)
		if err != nil {
			helpers.Error(c, 500, err.Error())
		} else {
			c.JSON(201, user)
		}
	}
}

func (u UserApi) Read(c *gin.Context) {
	users, err := u.controller.Read()
	if err != nil {
		helpers.Error(c, 500, err.Error())
	} else {
		c.JSON(200, users)
	}
}

func (u UserApi) Update(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	var user models.User
	json.Unmarshal(jsonData, &user)

	if err := u.Validate(user); err != nil {
		helpers.Error(c, 400, err.Error())
	} else {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			helpers.Error(c, 400, "Params Error: id (integer)")
		} else {
			if err := u.controller.Update(id, user); err != nil {
				if err.Error() == "record not found" {
					helpers.Error(c, 404, "User Not Found")
				} else {
					helpers.Error(c, 500, err.Error())
				}
			} else {
				c.JSON(200, gin.H{"Message": "Success"})
			}
		}
	}
}

func (u UserApi) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.Error(c, 400, "Params Error: id (integer)")
	} else {
		if err := u.controller.Delete(id); err != nil {
			if err.Error() == "record not found" {
				helpers.Error(c, 404, "User Not Found")
			} else {
				helpers.Error(c, 500, err.Error())
			}
		} else {
			c.JSON(200, gin.H{"Message": "Success"})
		}
	}
}
