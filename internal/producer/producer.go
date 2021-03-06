package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EgMeln/touchKafka/internal/config"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"strconv"
	time "time"
)

type Producer struct {
	Producer *kafka.Writer
}

func NewProducer(cfg *config.Config, url *string) *Producer {
	conn := kafka.Writer{
		Addr:         kafka.TCP(*url),
		Topic:        cfg.KafkaTopic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    2000,
		BatchTimeout: 70 * time.Microsecond,
	}
	return &Producer{Producer: &conn}
}

func (prod Producer) ProduceMessages(ctx context.Context) error {
	log.Printf("Start producing")
	t := time.Now()
	for i := 0; ; i++ {
		key := fmt.Sprintf("Key-%d", i)
		message, err := json.Marshal("this is message " + strconv.Itoa(i))
		if err != nil {
			return fmt.Errorf("marshal error")
		}
		msg := kafka.Message{
			Key:   []byte(key),
			Value: message,
		}
		err = prod.Producer.WriteMessages(ctx, msg)
		if err != nil {
			log.Info(err)
		} else {
			log.Info("produced", key)
		}
		if i%2000 == 0 {
			log.Info("send 2000 messages")
			log.Info(time.Since(t))
			t = time.Now()
		}
	}
}
