package middleware

import (
	"net/http"
	"strings"

	"project/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["type"] != "access" {
			c.JSON(401, gin.H{"error": "not access token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		c.Next()
	}
}

func RoleMiddleware(role string) gin.HandlerFunc {

	return func(c *gin.Context) {

		userRole, _ := c.Get("role")

		if userRole != role {
			c.JSON(403, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}