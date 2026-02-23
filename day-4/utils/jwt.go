package utils

import(
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var seceretKey = []byte("seceret_key")

func GenereteToken(userID uint, role string)(string, error){
	claims := jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(seceretKey)
}

func ValidateToken(tokenString string)(*jwt.Token, error){
	return jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		return seceretKey, nil
	})
}