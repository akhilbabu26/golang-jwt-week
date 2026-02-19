package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"

	"day-3/utils"
)

var users = make(map[string]string)

type SigninInput struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type RefereshInput struct{
	Refresh_token string `json:"refresh_token"`
}

func Signin(c *gin.Context){
	var signin SigninInput

	if err := c.ShouldBindJSON(&signin); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(signin.Password),
		bcrypt.DefaultCost,
	)
	if err != nil{
		c.JSON(500, gin.H{"error": "faild create hash pasword"})
		return
	}

	users[signin.Email] = string(hashPassword)
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "siginin success",
	})
}

func Login(c *gin.Context){
	var login LoginInput

	if err := c.ShouldBindJSON(&login); err != nil{
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	hashPassword, exists := users[login.Email]
	if !exists{
		c.JSON(400, gin.H{"error": "Your not a user"})
		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(hashPassword),
		[]byte(login.Password),
	)
	if err != nil{
		c.JSON(401, gin.H{"error": "invalid password"})
		return
	}

	accessToken, err := utils.GenereateAccessToken(1)
	if err != nil{
		c.JSON(500, gin.H{"error": "access token faild"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(1)
	if err != nil {
		c.JSON(500, gin.H{"error": "referesh token faild"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"refresh_token": refreshToken,
		"access_token": accessToken,
	})
}

func Refresh(c *gin.Context){
	var refreshToken RefereshInput

	if err := c.ShouldBindJSON(&refreshToken); err != nil{
		c.JSON(400, gin.H{"error": "Invalid referesh token"})
		return
	}

	token, err := utils.ValidateToken(refreshToken.Refresh_token)
	if err != nil || !token.Valid{
		c.JSON(401, gin.H{"error": "invalid token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	newAccessToken, _ := utils.GenereateAccessToken(userID)

	c.JSON(200, gin.H{
		"access_token": newAccessToken,
	})
}