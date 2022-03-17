package consumer

import (
	"context"
	"fmt"
	"github.com/EgMeln/touchKafka/internal/config"
	"github.com/EgMeln/touchKafka/internal/repository"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type Consumer struct {
	Consumer *kafka.Reader
}

func NewConsumer(cfg *config.Config, url *string) *Consumer {
	conn := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{*url},
		GroupID:  cfg.KafkaGroupID,
		Topic:    cfg.KafkaTopic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return &Consumer{Consumer: conn}
}
func (prod Consumer) ConsumeMessages(collection *repository.MongoRepository) {

	for {
		m, err := prod.Consumer.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		insertResult, err := collection.Collection.InsertOne(context.Background(), m)

		fmt.Println("Inserted : ", insertResult.InsertedID)
	}
}
