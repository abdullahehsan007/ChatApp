package mongodb

import (
	"chatapp/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *messageRepository) SendMessage(message model.Message, senderID string) (string, error) {
	// Create a MongoDB document
	msgDocument := bson.M{
		"senderId":    senderID,
		"receiverId":  message.Receiverid,
		"message":     message.Message,
		"timestamp":   time.Now(),
	}

	// Insert into MongoDB
	_, err := r.collection.InsertOne(context.Background(), msgDocument)
	if err != nil {
		return "", err
	}

	return "Message Sent Successfully", nil
}
