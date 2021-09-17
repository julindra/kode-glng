package middlewares

import "github.com/gin-gonic/gin"

func Error() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(500, gin.H{"error": err})
		}
	})
}
