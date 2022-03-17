package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	KafkaTopic   string `env:"KAFKATOPIC" envDefault:"myTopic"`
	KafkaGroupID string `env:"KafkaGID" envDefault:"myGroup"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("error with parsing env variables in kafka config %w", err)
	}
	return cfg, nil
}
