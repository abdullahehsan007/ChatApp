package mongodb

import (
	"chatapp/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *messageRepository) GetMessage(get model.Get, yourID string) ([]model.Message, error) {
	filter := bson.M{
		"senderId":   get.Senderid,
		"receiverId": yourID}

	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var rawMessages []bson.M
	cursor.All(context.Background(), &rawMessages)

	// Convert to simple format
	var messages []model.Message
	for _, raw := range rawMessages {
		msg := model.Message{
			Senderid:   raw["senderId"].(string),
			Receiverid: raw["receiverId"].(string),
			Message:    raw["message"].(string),
			Time:       raw["timestamp"].(primitive.DateTime).Time().Format("2006-01-02 15:04"),
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
