package kafka

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
	"strings"
)

type Producer struct {
	writer *kafka.Writer
	logger zerolog.Logger
}

func NewProducer(brokers string, logger zerolog.Logger) *Producer {
	kafkaBrokers := strings.Split(brokers, ";")
	return &Producer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  kafkaBrokers,
			Balancer: &kafka.LeastBytes{},
		}),
		logger: logger,
	}
}

func (p Producer) Publish(topic string, data []byte) error {
	return p.writer.WriteMessages(context.Background(),
		kafka.Message{
			Topic: topic,
			Value: data,
		},
	)
}

func (p Producer) Close() error {
	return p.writer.Close()
}
