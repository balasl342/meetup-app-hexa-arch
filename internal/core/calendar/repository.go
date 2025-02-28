package calendar

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBCalendarRepository struct {
	collection *mongo.Collection
}

func NewMongoDBCalendarRepository(client *mongo.Client, dbName, collectionName string) *MongoDBCalendarRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoDBCalendarRepository{collection: collection}
}

func (r *MongoDBCalendarRepository) SaveEvents(ctx context.Context, events []Event) error {
	var documents []interface{}
	for _, event := range events {
		documents = append(documents, event)
	}
	_, err := r.collection.InsertMany(ctx, documents)
	if err != nil {
		return fmt.Errorf("failed to insert events: %v", err)
	}
	return err
}

func (r *MongoDBCalendarRepository) GetAllEvents(ctx context.Context) ([]Event, error) {
	var events []Event // Define a slice to hold all the documents

	cursor, err := r.collection.Find(ctx, bson.M{}) // Find all documents
	if err != nil {
		return events, fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(ctx) // Ensure the cursor is closed after use

	// Decode all documents into the slice
	if err := cursor.All(ctx, &events); err != nil {
		return events, fmt.Errorf("failed to decode documents: %w", err)
	}
	return events, nil
}
