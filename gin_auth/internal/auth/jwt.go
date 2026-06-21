package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims

	Role string `json:"role"`
}

func CreateToken(jwtSecret string, userId string, role string) (string, error) {

	now := time.Now().UTC()
	expiry := now.Add(7 * 24 * time.Hour)

	claims := Claims{
		RegisteredClaims : jwt.RegisteredClaims{
			Subject: userId,
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiry),
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", fmt.Errorf("Signing of token failed : %w", err)
	}
	return signed, nil
}


func ParseAndValidateToken(jwtSecret string, tokenString string ) (Claims, error) {
	var claims Claims

	parsed, err := jwt.ParseWithClaims(tokenString, &claims,
	func (t *jwt.Token) (interface{}, error){
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("Unexpected signing method :%v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	},
	jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
)
    if err != nil {
		return Claims{}, fmt.Errorf("Parse token failed: %w", err)
	}

	if !parsed.Valid {
		return Claims{}, fmt.Errorf("Invalid token: %w", err)
	}
	if claims.Subject == " " {
		return Claims{}, errors.New("token missing subject")
	}

	return claims, nil
}