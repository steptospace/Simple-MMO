package auth

import (
	"Simple-MMO/Auth/config"
	"Simple-MMO/pg"
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

func Run(logger zerolog.Logger, cfg *config.Config) (*AuthData, error) {
	postgres, err := pg.Init(cfg.PostgresConfig, logger)
	if err != nil {
		logger.Error().Err(err)
	}

	auth := &AuthData{
		logger: logger,
	}
	auth.postgres = postgres
	return auth, nil
}

func (a *AuthData) tryAuth(updateInterval time.Duration, retryCount int) (login string, password string) {
	for i := 0; i < retryCount; i++ {
		//a.postgres.GetAccountData()
		time.Sleep(updateInterval)
	}
	return login, genHash(password)
}

func genHash(password string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(password))
	return hex.EncodeToString(algorithm.Sum(nil))
}
