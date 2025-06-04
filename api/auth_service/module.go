package authservice

import (
	mongodb "chatapp/db/MongoDB"
	postgresql "chatapp/db/Postgresql"
	"chatapp/model"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(ctx *gin.Context, user model.Info) error
	Login(email, password string) (string, error)
	Authenticator(credential model.User) (string, error)
	Authorize(token string) (bool, string, error)
	BearerToken(header string) string
	SendMessage(ctx *gin.Context, user model.Message, Token string) (string, error)
	GetMessage(ctx *gin.Context, get model.Get, Token string) (string, error)
}

type authService struct {
	repo  postgresql.UserRepository
	mongo mongodb.MessageRepository
}

func NewAuthService(repo postgresql.UserRepository) AuthService {
	return &authService{repo: repo}
}
