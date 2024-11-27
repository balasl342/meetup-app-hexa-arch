package auth

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	SecretKey string
}

func NewJWTService(secretKey string) *JWTService {
	return &JWTService{
		SecretKey: secretKey,
	}
}

func (j *JWTService) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		log.Println("Invalid token")
		return nil, err
	}
	return token, nil
}
