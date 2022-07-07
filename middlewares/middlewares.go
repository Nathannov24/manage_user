package middlewares

import (
	"fmt"
	"manage_user/constants"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
	SigningKey: []byte(constants.SECRET_JWT),
})

func CreateToken(userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires after 1 hour

	tokenString, err := token.SignedString([]byte(constants.SECRET_JWT))
	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func CurrentLoginUser(e echo.Context) int {
	token := e.Get("user").(*jwt.Token)
	if token != nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userId"]
		switch userId.(type) {
		case float64:
			return int(userId.(float64))
		default:
			return userId.(int)
		}
	}
	return -1 // invalid user
}

func ExtractTokenUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}
