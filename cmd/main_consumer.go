package main

import (
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
	mongoURL := os.Getenv("mongoURL")
	dbName := os.Getenv("dbName")
	collectionName := os.Getenv("collectionName")
	collection := repository.GetMongoCollection(mongoURL, dbName, collectionName)

	kafkaURL := os.Getenv("kafkaURL")
	cons := consumer.NewConsumer(cfg, &kafkaURL)

	if err != nil {
		log.Fatalf("error while creation producer %v", err)
	}
	defer func() {
		if err = cons.Consumer.Close(); err != nil {
			log.Fatalf("error closing 2 connection %v", err)
		}
	}()
	log.Println("consumer successfully created")
	cons.ConsumeMessages(collection)
	if err != nil {
		log.Fatalf("error while consuming messages - %e", err)
	}
	log.Println("successfully consume messages")
}
