package authservice

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey        = []byte(os.Getenv("JWT_Key"))
	refreshSecretKey = []byte(os.Getenv("JWT_REFRESH_KEY"))
)

func (s *authService) BearerToken(header string) string {
	if strings.HasPrefix(header, "Bearer ") {
		return strings.TrimPrefix(header, "Bearer ")
	}
	return ""
}

func (s *authService) ParseToken(token string, secret []byte) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}

func (s *authService) TokenType(token *jwt.Token, tokenType string) bool {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if tType, ok := claims["type"].(string); ok && tType == tokenType {
			return true
		}
	}
	return false
}

func (s *authService) Authorize(token string) (bool, string, error) {
	if token == "" {
		return false, "missing token", errors.New("authorization token required")
	}

	parsedToken, err := s.ParseToken(token, secretKey)
	if err == nil && s.TokenType(parsedToken, "access") {
		return true, "valid access token", nil
	}

	parsedToken, err = s.ParseToken(token, refreshSecretKey)
	if err == nil && s.TokenType(parsedToken, "refresh") {
		return true, "valid refresh token", nil
	}

	if err != nil && strings.Contains(err.Error(), "expired") {
		return false, "token expired", err
	}

	return false, "invalid token", errors.New("invalid token")
}