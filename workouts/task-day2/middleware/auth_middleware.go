package middlaware

import (
	"strings"

	"task-day2/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(401, gin.H{"error": "authorization header requored"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid{
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
