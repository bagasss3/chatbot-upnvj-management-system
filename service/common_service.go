package service

import (
	"cbupnvj/config"
	"cbupnvj/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateToken(userAuth *model.UserAuth, tokenType model.TokenType, tokenExpired time.Duration) (string, error) {
	expirationTime := time.Now().Add(tokenExpired)
	claims := &model.Claims{
		UserID: userAuth.UserID,
		Role:   userAuth.Role,
		Type:   tokenType,
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
