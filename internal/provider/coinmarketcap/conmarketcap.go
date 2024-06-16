package coinmarketcap

import (
	"currency-converter/pkg/logger"
	"net/http"
	"time"
)

type Provider struct {
	client *http.Client
	config Config
	logger logger.Logger
}

func New(config Config, logger logger.Logger) *Provider {
	return &Provider{
		client: &http.Client{
			Timeout: time.Second * config.TimeoutSec,
		},
		config: config,
		logger: logger,
	}
}
