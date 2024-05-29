package middleware

import (
	"PongPedia/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	mid "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

var IsLoggedIn = mid.JWTWithConfig(mid.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SCREAT_JWT),
})

// Create Token Jwt
func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["role_type"] = role
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SCREAT_JWT))
}

func IsAdmin(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["role_type"] != constants.ADMIN {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	userId := int(claims["user_id"].(float64))

	return userId, nil
}

func IsUser(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["role_type"] != constants.PLAYER {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	userId := int(claims["user_id"].(float64))

	return userId, nil
}
