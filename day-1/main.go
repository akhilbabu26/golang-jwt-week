package main

import (
	"day-1/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUsers)

	r.Run() // http://localhost:8080
}