package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

func GetMongoCollection(ctx context.Context, mongoURL, dbName, collectionName string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, fmt.Errorf("mongo connect error %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("client ping error %v", err)
	}
	fmt.Println("Connected to MongoDB ... !!")
	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	return &MongoRepository{
		Collection: collection,
	}, nil
}
