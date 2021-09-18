package middlewares

import (
	"encoding/json"
	"io/ioutil"

	"session7/models"

	"github.com/gin-gonic/gin"
)

func Parser(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	var order models.Order
	json.Unmarshal(jsonData, &order)
	c.Set("order", order)
	c.Next()
}
