package messageservice

import (
	"chatapp/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (s *messageService) SendMessage(ctx *gin.Context, user model.Message, Token string) (string, error) {
	id, err := DecodeToken(Token)
	if err != nil {
		return "", err
	} else {
		user.Senderid = id
		remarks, err := s.mongorepo.SendMessage(user, id)
		if err != nil {
			return "", err
		} else {
			return remarks, nil
		}
	}

}

func DecodeToken(tokenString string) (string, error) {
	// Parse the token without verification first
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", fmt.Errorf("invalid token: %v", err)
	}

	// Type assert the claims to jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims format")
	}

	// Extract the user ID from claims
	idClaim, exists := claims["id"]
	if !exists {
		return "", fmt.Errorf("id claim not found in token")
	}

	// Convert the ID to string
	id, ok := idClaim.(string)
	if !ok {
		return "", fmt.Errorf("id claim is not a string")
	}

	return id, nil
}
