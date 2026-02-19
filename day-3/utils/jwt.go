package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secertKey = []byte("seceret_key")

func GenereateAccessToken(userID uint)(string, error){
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(1 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secertKey)
}

func GenerateRefreshToken(userID uint)(string, error){
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secertKey)
}

func ValidateToken(tokenString string)(*jwt.Token, error){
	return jwt.Parse(tokenString, func (token *jwt.Token)(interface{}, error){
		return secertKey, nil
	})
}