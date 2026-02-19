package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"day-2/utils"
)

var users = make(map[string]string)

//post  /signup
func Signup(c *gin.Context){
	var input struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	users[input.Email] = string(hashedPassword)

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

// post /login
func Login(c *gin.Context){
	var input struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	c.BindJSON(&input)

	hashedpassword, exists := users[input.Email]
	if !exists{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedpassword),
		[]byte(input.Password),
	)

	if err != nil {
		c.JSON(401, gin.H{"error": "invalid credential"})
		return
	}

	accessToken, err := utils.GenerateAccessToken(1)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate access token"})
		return
	}

	refereshToken, err := utils.GenerateRefereshToken(1)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate refresh token"})
		return
	}

	c.JSON(200, gin.H{
		"access_token": accessToken,
		"referesh_token": refereshToken,
	})
}

//post /refresh
func Referesh(c *gin.Context){
	var body struct{
		RefereshToken string `json:"refresh_token"`
	}

	c.BindJSON(&body)

	token, err := jwt.Parse(body.RefereshToken, func(token *jwt.Token) (interface{}, error){
		return []byte("super_secret_key"), nil
	})

	if err != nil || !token.Valid{
		c.JSON(401, gin.H{"error": "invalid referesh token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	newAccessToken, _ := utils.GenerateAccessToken(userID)

	c.JSON(200, gin.H{"access_token": newAccessToken})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Logged out"})
}
