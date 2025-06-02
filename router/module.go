package router

import (
	authservice "chatapp/api/auth_service"
	"chatapp/model"

	"github.com/gin-gonic/gin"
)

type Router interface {
	RoutersSetup(ctx *gin.Engine)
}

type Service interface {
	SignUp(ctx *gin.Context, user model.Info) error
	Login(email, password string) (string, error)
	SendMessage(ctx *gin.Context, user model.Message, Token string) (string, error)
}

type routerImpl struct {
	service     Service
	authService authservice.AuthService
}

func NewRouter(service Service, authService authservice.AuthService) Router {
	return &routerImpl{
		service:     service,
		authService: authService,
	}
}
