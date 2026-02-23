package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"day-3/config"
	"day-3/modules"
	"day-3/utils"
)

func Signup(c *gin.Context){
	var user modules.Users

	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashed,_ := utils.HashPassword(user.Password)
	user.Password = hashed

	if user.Role == ""{
		user.Role = "user"
	}

	config.DB.Create(&user)

	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context){
	var loginInput modules.Users
	var user modules.Users

	c.ShouldBindJSON(&loginInput)

	config.DB.Where("email = ?", loginInput.Email).First(&user)

	if !utils.CheckPassword(loginInput.Password, user.Password){
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, _ := utils.GenereteToken(user.ID, user.Role)

	c.JSON(200, gin.H{
		"access_token": token,
	})

}