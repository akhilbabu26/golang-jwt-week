package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
}

var users []User

func main(){
	r := gin.Default()

	r.POST("/signup", Signup)
	r.POST("/login", Login)

	r.Run(":8080")
}

func Signup(c *gin.Context){
	var body User

	c.ShouldBindJSON(&body)

	users = append(users, body)

	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}

func Login(c *gin.Context){
	var body User

	c.ShouldBindJSON(&body)

	for _, user := range users{
		if user.Email == body.Email{
			c.JSON(http.StatusOK, gin.H{"welcome": user.Name})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error":"user not found"})
}