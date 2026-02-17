package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Users struct{
	Name string `json:"name" binding:"required"`
}

var users = []Users{
	{Name: "Akhil"},
	{Name: "Junaid"},
	{Name: "Nihal"},
}

func GetUsers(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func AddUser(c *gin.Context){
	var user Users

	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	users = append(users, user)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}