package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"project/config"
	"project/models"
	"project/utils"
)

func Signup(c *gin.Context){ //
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := utils.HashPassword(body.Password)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "cant hash password"})
		return
	}

	if body.Role == ""{
		body.Role = "user"
	}

	user := models.User{
		Name: body.Name,
		Email: body.Email,
		Password: hash,
		Role: body.Role,
	}

	config.DB.Create(&user)
	c.JSON(200, gin.H{"message": "user created"})
}

func Login(c * gin.Context){
	var body models.User

	c.ShouldBindJSON(&body)

	var user models.User
	config.DB.Where("email = ?", body.Email).First(&user)

	if !utils.CheckPassword(body.Password, user.Password){
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	accessToken, _ := utils.AccessToken(user.ID, user.Role)
	refreshToken, _ := utils.RefereshToken(user.ID, user.Role)

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(c *gin.Context) {

	var body struct {
		RefreshToken string `json:"refresh_token"`
	}

	c.ShouldBindJSON(&body)

	token, err := utils.ValidateToken(body.RefreshToken)
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "invalid refresh token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	if claims["type"] != "refresh" {
		c.JSON(401, gin.H{"error": "wrong token type"})
		return
	}

	userID := uint(claims["user_id"].(float64))

	var user models.User
	config.DB.First(&user, userID)

	newAccessToken, _ := utils.AccessToken(user.ID, user.Role)

	c.JSON(200, gin.H{
		"access_token": newAccessToken,
	})
}