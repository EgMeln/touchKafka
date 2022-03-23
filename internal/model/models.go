package model

type KafkaMessage struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}
