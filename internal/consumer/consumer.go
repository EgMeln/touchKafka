package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EgMeln/touchKafka/internal/config"
	"github.com/EgMeln/touchKafka/internal/model"
	"github.com/EgMeln/touchKafka/internal/repository"
	"github.com/jackc/pgx/v4"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type Consumer struct {
	Consumer *kafka.Conn
}

func NewConsumer(cfg *config.Config, url *string) (*Consumer, error) {
	//conn := kafka.NewReader(kafka.ReaderConfig{
	//	Brokers:  []string{*url},
	//	GroupID:  cfg.KafkaGroupID,
	//	Topic:    cfg.KafkaTopic,
	//	MinBytes: 10e3, // 10KB
	//	MaxBytes: 10e6, // 10MB
	//})
	conn, err := kafka.DialLeader(context.Background(), "tcp", *url, cfg.KafkaTopic, 0)
	if err != nil {
		return nil, fmt.Errorf("can't create new consumer %v", err)
	}
	return &Consumer{Consumer: conn}, nil
}
func (cons Consumer) ConsumeMessages(connection *repository.PostgresConnection) error {
	pgxBatch := &pgx.Batch{}
	batch := cons.Consumer.ReadBatch(10e3, 10e6)
	count := 0
	t := time.Now()
	for {
		m, err := batch.ReadMessage()
		if err != nil {
			log.Info("BREAK")
			break
		}
		message := model.KafkaMessage{}
		message.Key = string(m.Key)
		err = json.Unmarshal(m.Value, &message.Message)
		if err != nil {
			return fmt.Errorf("can't parse message")
		}
		log.Info(message.Key)
		log.Info(message.Message)
		pgxBatch.Queue("insert into kafka(key, message) values($1, $2)", message.Key, message.Message)
		count++
		fmt.Println("Inserted : ", message)
		if count%2000 == 0 {
			batchResult := connection.Conn.SendBatch(context.Background(), pgxBatch)
			ct, err := batchResult.Exec()
			if err != nil {
				return fmt.Errorf("can't insert messages into repository %v", err)
			}
			if ct.RowsAffected() != 1 {
				return fmt.Errorf("RowsAffected() => %v, want %v", ct.RowsAffected(), 1)
			}
			log.Info("send 2000 messages")
			pgxBatch = &pgx.Batch{}
			log.Info(time.Since(t))
			t = time.Now()
			time.Sleep(1 * time.Second)
		}
	}
	if err := batch.Close(); err != nil {
		return fmt.Errorf("error while closing batch %v", err)
	}
	log.Info(time.Since(t))
	t = time.Now()
	return nil
}
