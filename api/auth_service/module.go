package authservice

import (
	"chatapp/db"
	"chatapp/model"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(ctx *gin.Context, user model.Info) error
	Login(email, password string) (string, error)
	Authenticator(credential model.User) (string, error)
	Authorize(token string) (bool, string, error)
	BearerToken(header string) string
}

type authService struct {
	repo db.UserRepository
}

func NewAuthService(repo db.UserRepository) AuthService {
	return &authService{repo: repo}
}
