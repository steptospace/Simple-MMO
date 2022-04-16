package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	KafkaBrokers   string `split_words:"true" default:"localhost:29092"` // ; separated brokers
	KafkaTopicOut  string `split_words:"true" default:"simpleMMO.data"`
	PrometheusPort string `split_words:"true" default:"3737"`

	PostgresConfig string        `split_words:"true" default:"host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"`
	RetryCount     int           `split_words:"true" default:"5"`
	UpdateInterval time.Duration `split_word:"true" default:"1m"`
}

func Init() (*Config, error) {
	var cfg Config

	err := envconfig.Process("AUTH", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
