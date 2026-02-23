package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"day-3/config"
	"day-3/modules"
)

func GetUser(c *gin.Context){
	var users []modules.Users
	config.DB.Find(&users)

	c.JSON(200, users)
}


func CreateUser(c *gin.Context){
	var user modules.Users
	c.ShouldBindJSON(&user)

	config.DB.Create(&user)
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context){
	id  := c.Param("id")

	var user modules.Users
	config.DB.First(&user, id)

	var updateInput modules.Users
	c.ShouldBindJSON(&updateInput)

	user.Name = updateInput.Name
	user.Email = updateInput.Email
	user.Role = updateInput.Role

	config.DB.Save(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context){
	id := c.Param("id")

	config.DB.Delete(&modules.Users{}, id)

	c.JSON(http.StatusOK, gin.H{"message": "delete"})
}