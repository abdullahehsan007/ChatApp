package router

import (
	authservice "chatapp/api/auth_service"
	messageservice "chatapp/api/message_service"

	"github.com/gin-gonic/gin"
)

type Router interface {
	RoutersSetup(ctx *gin.Engine)
}

type Service interface {
	handleSignUp(ctx *gin.Context)
	Login() gin.HandlerFunc
	Refresh() gin.HandlerFunc
	Authorize() gin.HandlerFunc
	SendMessage() gin.HandlerFunc
	GetMessage() gin.HandlerFunc
}

type routerImpl struct {
	service Service
}

type serviceImpl struct {
	authService    authservice.AuthService
	messageService messageservice.MessageService
}

func NewRouter(service Service) Router {
	return &routerImpl{
		service: service,
	}
}

func NewService(authService authservice.AuthService, messageService messageservice.MessageService) Service {
	return &serviceImpl{
		authService:    authService,
		messageService: messageService,
	}
}
