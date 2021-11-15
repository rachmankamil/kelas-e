package middleware

import (
	"km-kelas-e/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID int) (string, error) {
	//TODO : check DB for rules or any other data that necessary

	payload := jwt.MapClaims{
		"userid": userID,
		"role":   "admin",
		"exp":    time.Now().Add(time.Hour * 6).Unix(),
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)
	token, error := tokenWithClaims.SignedString([]byte(config.JwtSecret))
	return token, error
}

func ExtractClaim(e echo.Context) (claims map[string]interface{}) {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims = user.Claims.(jwt.MapClaims)
	}

	return
}
