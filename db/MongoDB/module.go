package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MessageRepository struct {
	collection *mongo.Collection
}

func NewMessageRepository(db *mongo.Database, collectionName string) *MessageRepository {
	return &MessageRepository{
		collection: db.Collection(collectionName),
	}
}
