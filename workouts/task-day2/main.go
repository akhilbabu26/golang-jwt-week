package main

import (
	"github.com/gin-gonic/gin"

	"task-day2/controllers"
	"task-day2/middleware"
)

func main(){
	r := gin.Default()

	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	r.POST("/refresh", controller.Refresh)
	r.POST("/logout", controller.Logout)

	private := r.Group("/private")
	private.Use(middlaware.AuthMiddleware())
	private.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to private route",
		})
	})

	r.Run(":8080")
}