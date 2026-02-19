package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret_key")

func GenerateAccessToken(userID uint)(string, error){
	claims := jwt.MapClaims{  // creates claim that you pass through token
		"user_id": userID,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // creates new token with signature and claim
	return token.SignedString(secretKey)
}

func GenerateRefreshToken(userID uint)(string, error){
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string)(*jwt.Token, error){
	return jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		return secretKey, nil
	})
}