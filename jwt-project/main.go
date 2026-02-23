package main

import (
	"project/config"
	"project/controllers"
	"project/middleware"
	"project/models"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectBD()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/refresh", controllers.RefreshToken)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "logged in user"})
	})

	admin := auth.Group("/")
	admin.Use(middleware.RoleMiddleware("admin"))

	admin.GET("/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "admin route"})
	})

	r.Run(":8080")
}