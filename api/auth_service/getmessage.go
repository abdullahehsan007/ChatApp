package authservice

import (
	"chatapp/model"

	"github.com/gin-gonic/gin"
)

func (s *authService) GetMessage(ctx *gin.Context, get model.Get, Token string) (string, error) {
	id, err := DecodeToken(Token)
	// if err != nil {
	// 	return "", err
	// } else {
	// 	get.Receiverid = id
	// 	remarks, err := s.repo.(model.Message,id)
	// 	if err != nil {
	// 		return "", err
	// 	} else {
	// 		return remarks, nil
	// 	}
	// }
	return id, err
}
