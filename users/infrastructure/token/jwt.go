package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func Generate(sub string) (string, error) {
	claims := jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
}
