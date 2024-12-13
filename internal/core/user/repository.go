package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

// Save saves a new user to MongoDB.
func (r *MongoDBUserRepository) Save(ctx context.Context, user *User) error {
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// GetByEmail retrieves a user from MongoDB by their email.
func (r *MongoDBUserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	filter := bson.D{{Key: "email", Value: email}}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
