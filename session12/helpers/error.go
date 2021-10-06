package helpers

import (
	"session12/models"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, code int, error string) {
	c.JSON(code, &models.Error{Error: error})
}
