package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"task-day2/utils"
)

var users = make(map[string]string) // for email and hashed password

// post /signup
func Signup(c *gin.Context){
	var input struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if err != nil{
		c.JSON(500, gin.H{"error": "failed to hash password"})
		return
	}

	users[input.Email] = string(hashedPassword)

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

//post /login
func Login(c *gin.Context){
	var input struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, exists := users[input.Email]
	if !exists{
		c.JSON(401, gin.H{"error": "invalid email"})
		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(input.Password),
	)
	if err != nil{
		c.JSON(500, gin.H{"error": "invalid password"})
	}

	accessToken, err := utils.GenerateAccessToken(1) // 1 is fixed id for study
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate access token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(1)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate refresh token"})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

// post /refresh
func Refresh(c *gin.Context){
	var body struct{
		RefereshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&body); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.ValidateToken(body.RefereshToken)
	if err != nil || !token.Valid{
		c.JSON(401, gin.H{"error": "invalid referesh token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	newAccessToken, _ := utils.GenerateAccessToken(userID)

	c.JSON(200, gin.H{
		"access_token": newAccessToken,
	})
}

// POST /logout
func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logged out",
	})
}