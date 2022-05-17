package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId float64, name, email string) (string, error) {
	claims := jwt.MapClaims{}

	claims["userId"] = userId
	claims["name"] = name

	claims["expired"] = time.Now().Add(time.Hour * 3).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("TODO"))
}

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return userId
	}
	return 0
}

func ExtractTokenName(e echo.Context) string {
	name := e.Get("user").(*jwt.Token)
	if name.Valid {
		claims := name.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return name
	}
	return ""
}
