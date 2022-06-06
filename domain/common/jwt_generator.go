package common

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken(userid, name string) (string, error) {
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")

	// Set expires minutes count for secret key from .env file.
	daysCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_DAYS_COUNT"))

	// Create a new claims.
	// Create the Claims
	claims := jwt.MapClaims{
		"id":    userid,
		"name":  name,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * time.Duration(daysCount)).Unix(),
	}

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
