package authservice

import (
	"chatapp/model"

	"github.com/gin-gonic/gin"
)

func (s *authService) SendMessage(ctx *gin.Context, user model.Message) error {
	exists, err := s.repo.GetUserByID(user.Senderid)
	if err != nil {
		return err
	}
	if exists {
		return s.repo.SendMessage(user)
	}

}
