package producer

import (
	"context"
	"fmt"
	"github.com/EgMeln/touchKafka/internal/config"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type Producer struct {
	Producer *kafka.Writer
}

func NewProducer(cfg *config.Config, url *string) *Producer {
	conn := kafka.Writer{
		Addr:     kafka.TCP(*url),
		Topic:    cfg.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{Producer: &conn}
}

func (prod Producer) ProduceMessages() {
	//var messages []kafka.Message
	//for i := 0; i < 2000; i++ {
	//	message := model.Message{
	//		Value: strconv.Itoa(i),
	//		Time:  time.Now(),
	//	}
	//	msg, err := json.Marshal(message)
	//	if err != nil {
	//		return fmt.Errorf("can't marshal message %w", err)
	//	}
	//	messages = append(messages, kafka.Message{Value: msg})
	//	if i == 1999 {
	//		log.Info("send 2000 messages")
	//		time.Sleep(5 * time.Second)
	//	}
	//}
	//err := prod.Producer.SetWriteDeadline(time.Now().Add(10 * time.Second))
	//if err != nil {
	//	log.Fatal("failed ser write deadline:", err)
	//}
	//_, err = prod.Producer.WriteMessages(messages...)
	//if err != nil {
	//	log.Fatal("can't send messages:", err)
	//}
	log.Printf("Start producing")
	for i := 0; ; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprint(uuid.New())),
		}
		err := prod.Producer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("produced", key)
		}
		if i == 1999 {
			log.Info("send 2000 messages")
			time.Sleep(5 * time.Second)
		}
	}
}
