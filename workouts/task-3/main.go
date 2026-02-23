package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
}

func main(){
	r := gin.Default()

	r.GET("/users", ListUser)

	r.Run(":8080")
}

func ListUser(c *gin.Context){
	body := []User{
		{Name: "Akhil", Email: "akhil@gmail.com"},
		{Name: "John", Email: "john@gmail.com"},
		{Name: "Sara", Email: "sara@gmail.com"},
	}

	c.JSON(http.StatusOK, gin.H{
		"users": body,
	})
}