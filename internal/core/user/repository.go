package user

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBUserRepository is an implementation of the UserRepository interface using MongoDB.
type MongoDBUserRepository struct {
	collection *mongo.Collection
}

// NewMongoDBUserRepository creates a new MongoDB-backed user repository.
func NewMongoDBUserRepository(client *mongo.Client, dbName, collectionName string) *MongoDBUserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoDBUserRepository{collection: collection}
}
