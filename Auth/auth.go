package auth

import (
	"Simple-MMO/Auth/config"
	"Simple-MMO/kafka"
	"Simple-MMO/pg"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/rs/zerolog"
	"time"
)

type AuthData struct {
	logger   zerolog.Logger
	login    string
	password string
	postgres *pg.Worker
}

func Auth(logger zerolog.Logger, cfg *config.Config, ctx context.Context) (<-chan kafka.UserData, error) {
	solution := make(chan kafka.UserData)
	// Add kafka and metrics
	authData, err := Run(logger, cfg)
	if err != nil {

		return nil, err
	}
	authData.tryAuth(cfg.UpdateInterval, cfg.RetryCount)

	//Сбор данных об истории мира

	return solution, nil
}

func Run(logger zerolog.Logger, cfg *config.Config) (*AuthData, error) {
	postgres, err := pg.Init(cfg.PostgresConfig, logger)
	if err != nil {
		logger.Error().Err(err)
	}

	//test
	postgres.InsertAccount("test", "test")

	auth := &AuthData{
		logger: logger,
	}
	auth.postgres = postgres
	return auth, nil
}

/// Use ticker
func (a *AuthData) tryAuth(updateInterval time.Duration, retryCount int) (login string, password string) {

	for i := 0; i <= retryCount; i++ {
		time.Sleep(updateInterval)
		//a.postgres.GetAccountData()
	}

	return login, genHash(password)
}

func genHash(password string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(password))
	return hex.EncodeToString(algorithm.Sum(nil))
}
