package main

import (
	"github.com/gin-gonic/gin"

	"day-2/controllers"
	"day-2/middleware"
)

func main(){
	r := gin.Default()

	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	r.POST("/refresh", controller.Referesh)
	r.POST("/logout", controller.Logout)

	private := r.Group("/private")

	private.Use(middleware.AuthMiddleware())
	private.GET("/profile", func(c *gin.Context){
		c.JSON(200, gin.H{"message": "Welcome to private route"})
	})

	r.Run(":8080")
}