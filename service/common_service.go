package service

import (
	"cbupnvj/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID int64 `json:"userID"`
	jwt.StandardClaims
}

func generateToken(id int64, tokenExpired time.Duration) (string, error) {
	expirationTime := time.Now().Add(tokenExpired)
	claims := &Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := createToken.SignedString([]byte(config.JWTKey()))
	if err != nil {
		return "", err
	}
	return token, nil
}
