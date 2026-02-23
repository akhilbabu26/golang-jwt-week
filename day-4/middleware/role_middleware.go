package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc{
	return func (c *gin.Context){
		role, exists := c.Get("role")

		if !exists{
			c.JSON(http.StatusForbidden, gin.H{"error": "role not found"})
			c.Abort()
			return
		}

		if role != "admin"{
			c.JSON(http.StatusForbidden, gin.H{"error": "only admin can enter"})
			c.Abort()
			return
		}

		c.Next()
	}
}