package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct{
	Name string `json:"name" binding:"required"`
}
var users = []User{
	{Name: "Akhiil"},
	{Name: "Nihal"},
}

func GetUsers(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func CreateUsers(c *gin.Context){
	var user User

	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users = append(users, user)

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}