package main

import (
	"context"
	"github.com/EgMeln/touchKafka/internal/config"
	"github.com/EgMeln/touchKafka/internal/consumer"
	"github.com/EgMeln/touchKafka/internal/repository"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln("Config error: ", cfg)
	}
	//mongoURL := os.Getenv("mongoURL")
	//dbName := os.Getenv("dbName")
	//collectionName := os.Getenv("collectionName")
	//collection := repository.GetMongoCollection(mongoURL, dbName, collectionName)
	ctx := context.Background()
	kafkaURL := os.Getenv("kafkaURL")
	cons, err := consumer.NewConsumer(ctx, cfg, &kafkaURL)
	if err != nil {
		log.Fatalf("error while creation consumer %v", err)
	}
	dbNamePostgres := os.Getenv("postgresURL")
	conn, err := repository.GetPostgresConnection(ctx, dbNamePostgres)
	if err != nil {
		log.Fatalf("postgres connection error %v", err)
	}
	defer func() {
		if err = cons.Consumer.Close(); err != nil {
			log.Fatalf("error closing connection %v", err)
		}
	}()
	log.Println("consumer successfully created")
	//cons.ConsumeMessages(collection)

	err = cons.ConsumeMessages(ctx, conn)
	if err != nil {
		log.Fatalf("error while consuming messages - %e", err)
		return
	}
	log.Println("successfully consume messages")
}
