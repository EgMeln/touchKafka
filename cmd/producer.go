package main

import (
	"github.com/EgMeln/touchKafka/internal/config"
	"github.com/EgMeln/touchKafka/internal/producer"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln("Config error: ", cfg)
	}
	kafkaURL := os.Getenv("kafkaURL")
	prod := producer.NewProducer(cfg, &kafkaURL)
	if err != nil {
		log.Fatalf("error while creation producer %v", err)
	}
	defer func() {
		if err = prod.Producer.Close(); err != nil {
			log.Fatalf("error closing 2 connection %v", err)
		}
	}()
	log.Println("producer successfully created")

	prod.ProduceMessages()
	if err != nil {
		log.Fatalf("error while sending messages - %e", err)
	}
	log.Println("successfully send messages")

}
