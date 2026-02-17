package main

import (
	"project-mod/controllers"
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()

	r.GET("/users", controller.GetUsers)
	r.POST("/users", controller.AddUser)

	r.Run(":8080")
}
