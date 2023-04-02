package middleware

import (
	"cbupnvj/config"
	"cbupnvj/model"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		authHeader := c.Request().Header.Get("Authorization")
// 		if authHeader == "" {
// 			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing Authorization header"})
// 		}

// 		tokenString := strings.Split(authHeader, " ")[1]
// 		claims := &model.Claims{}

// 		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid token")
// 			}
// 			return config.JWTKey(), nil
// 		})

// 		if err != nil {
// 			if err == jwt.ErrSignatureInvalid {
// 				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token signature"})
// 			}
// 			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
// 		}

// 		if !token.Valid {
// 			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
// 		}

// 		c.Set("userID", claims.UserID)
// 		return next(c)
// 	}
// }

func MustAuthenticateAccessToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing Authorization header"})
			}

			tokenString := strings.Split(authHeader, " ")[1]
			claims := &model.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					logrus.Error(ok)
					return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid token")
				}
				return []byte(config.JWTKey()), nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					logrus.Error(err)
					return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token signature"})
				}
				if jwt.ValidationErrorExpired != 0 {
					logrus.Error(err)
					return c.JSON(http.StatusForbidden, echo.Map{"error": "token has expired"})
				}
				logrus.Error(err)
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
			}

			if !token.Valid {
				logrus.Error(token.Valid)
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
			}

			c.Set("userID", claims.UserID)
			c.Set("userRole", claims.Role)
			return next(c)
		}
	}
}
