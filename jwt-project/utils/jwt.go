package utils

import(
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var seceretKey = []byte("seceret_key")

func AccessToken(userID uint, role string)(string, error){
	claims := jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"type": "access",
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(seceretKey)
}

func RefereshToken(userID uint, role string)(string, error){
	claims := jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"type": "access",
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(seceretKey)
}

func ValidateToken(tokenString string)(*jwt.Token, error){
	return jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		return seceretKey, nil
	})
}