package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JWTSecretKey = []byte("jwt-secret-keys")

func GetJWTToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"iss": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ClaimJWTUserId(tokenString interface{}) int {
	jwtToken := tokenString.(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	userId := int(claims["sub"].(float64))

	return userId
}
