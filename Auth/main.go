package auth

import (
	"Simple-MMO/Auth/config"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	config, err := config.Init()
	if err != nil {
		logger.Error().Err(err).Msg("Error initializing config")
		return
	}
	// Add kafka and metrics
	authData, err := Run(logger, config)
	if err != nil {
		logger.Error().Err(err)
	}

	authData.tryAuth(config.UpdateInterval, config.RetryCount)
}
