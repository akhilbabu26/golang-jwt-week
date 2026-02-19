package main

import (
	"github.com/gin-gonic/gin"

	"day-3/controllers"
)

func main(){
	r := gin.Default()

	r.POST("/signin", controller.Signin)
	r.POST("/login", controller.Login)
	r.POST("/refresh", controller.Refresh)
	
	r.Run(":8080")

}