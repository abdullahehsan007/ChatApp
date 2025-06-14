package mongodb

import (
	"chatapp/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository interface {
	SendMessage(message model.Message, senderID string) (string, error)
	GetMessage(get model.Get, yourID string) ([]model.Message, error)
}

type messageRepository struct {
	collection *mongo.Collection
}

func NewMessageRepository(collection *mongo.Collection) MessageRepository {
	return &messageRepository{collection: collection}
}
