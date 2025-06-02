package authservice

import (
	"chatapp/model"
	"errors"

	"github.com/gin-gonic/gin"
)
func (s *authService) SignUp(ctx *gin.Context, user model.Info) error {
	exists, err := s.repo.GetUser(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}
	return s.repo.CreateUser(user)
}
