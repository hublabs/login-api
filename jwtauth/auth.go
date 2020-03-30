package jwtauth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	jwtSigningMethod = jwt.SigningMethodHS256
	JwtSecret        = "JWT_SECRET"
)

const (
	JWT_KEY = "JWT_SECRET"
)

func init() {
	if s := os.Getenv(JWT_KEY); s != "" {
		JwtSecret = s
	}
}
func NewToken(m map[string]interface{}, auth string) (string, error) {
	claims := jwt.MapClaims{
		"iss": "colleague",
		"aud": auth,
		"nbf": time.Now().Add(-time.Minute * 5).Unix(),
		"exp": time.Now().Add(time.Hour * 15).Unix(),
	}
	for k, v := range m {
		claims[k] = v
	}
	return jwt.NewWithClaims(jwtSigningMethod, claims).SignedString([]byte(JwtSecret))
}
