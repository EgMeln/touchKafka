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
	//mongoURL := os.Getenv("mongoURL")
	//dbName := os.Getenv("dbName")
	//collectionName := os.Getenv("collectionName")
	//collection := repository.GetMongoCollection(mongoURL, dbName, collectionName)

	kafkaURL := os.Getenv("kafkaURL")
	cons, err := consumer.NewConsumer(cfg, &kafkaURL)
	if err != nil {
		log.Fatalf("error while creation consumer %v", err)
	}
	dbNamePostgres := os.Getenv("postgresURL")
	conn := repository.GetPostgresConnection(dbNamePostgres)
	defer func() {
		if err = cons.Consumer.Close(); err != nil {
			log.Fatalf("error closing 2 connection %v", err)
		}
	}()
	log.Println("consumer successfully created")
	//cons.ConsumeMessages(collection)

	err = cons.ConsumeMessages(conn)
	if err != nil {
		log.Fatalf("error while consuming messages - %e", err)
		return
	}
	log.Println("successfully consume messages")
}
