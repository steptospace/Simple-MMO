package main

import (
	auth "Simple-MMO/Auth"
	"Simple-MMO/Auth/config"
	"Simple-MMO/kafka"
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	config, err := config.Init()
	if err != nil {
		logger.Error().Err(err)
	}
	producer := kafka.NewProducer(config.KafkaBrokers, logger)
	logger.Info().Str("brokers", config.KafkaBrokers).Msg("Kafka producer created")
	ctx, stopContext := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// Сделать пока что по тупому
	go func() {
		runChan, err := auth.Auth(logger, config, ctx)
		if err != nil {
			logger.Error().Err(err)
			stopContext()
			return
		}
	running:
		for {
			select {
			case data := <-runChan:
				fmt.Println(data, producer)
				// Обращение к функциям движка
			case <-ctx.Done():
				break running
			}
		}
	}()

}
