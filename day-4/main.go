package main

import(
	"day-3/config"
	"day-3/modules"
	"day-3/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	config.ConnectDB()

	config.DB.AutoMigrate(&modules.Users{})

	routes.SetupRoutes(r)

	r.Run(":8080")
}