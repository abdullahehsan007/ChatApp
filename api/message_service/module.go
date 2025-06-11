package messageservice

import (
	mongodb "chatapp/db/MongoDB"
	"chatapp/model"

	"github.com/gin-gonic/gin"
)

type MessageService interface {
	SendMessage(ctx *gin.Context, user model.Message, Token string) (string, error)
	// GetMessage(get model.Get, Token string) (string, error)
}

type messageService struct {
	mongorepo mongodb.MessageRepository
}

func NewMessageService(mongorepo mongodb.MessageRepository) MessageService {
	return &messageService{
		mongorepo: mongorepo,
	}
}
