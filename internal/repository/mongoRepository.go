package repository

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

func GetMongoCollection(mongoURL, dbName, collectionName string) *MongoRepository {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB ... !!")
	db := client.Database(dbName)
	collection := db.Collection(collectionName)
	return &MongoRepository{
		Collection: collection,
	}
}
