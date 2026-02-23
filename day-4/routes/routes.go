package routes

import(
	"github.com/gin-gonic/gin"
	"day-3/controllers"
	"day-3/middleware"
)

func SetupRoutes(r *gin.Engine){
	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())

	admin.GET("/users", controller.GetUser)
	admin.POST("/users", controller.CreateUser)
	admin.PUT("/users/:id", controller.UpdateUser)
	admin.DELETE("/users/:id", controller.DeleteUser)
}