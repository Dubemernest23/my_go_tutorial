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
	exp := now.Add(7 * 24 * time.Hour)

	claim := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		Role: role,
	}

	// ✅ HS256 matches []byte(jwtSecret)
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signed, err := tok.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("Failed to sign token: %w", err)
	}
	return signed, nil
}

func ParseToken(jwtSecret string, tokenStr string) (Claims, error) {
	var claims Claims

	parsed, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil

	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return Claims{}, fmt.Errorf("Failed to parse token: %w", err)
	}

	if !parsed.Valid {
		return Claims{}, fmt.Errorf("Invalid token")
	}

	if claims.Subject == "" {
		return Claims{}, errors.New("token missing subject")
	}

	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now().UTC()) {
		return Claims{}, errors.New("token expired")
	}

	return claims, nil
}
