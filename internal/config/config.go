package config

import (
	"currency-rates/internal/provider/coinmarketcap"
	"currency-rates/pkg/logger"
	"github.com/pelletier/go-toml/v2"
	"os"
	"sync"
)

const (
	configFileName = "config.toml"
)

type Config struct {
	LogLevel      string               `toml:"logLevel"`
	CoinMarketCap coinmarketcap.Config `toml:"CoinMarketCap"`
}

var (
	config Config
	once   sync.Once
)

func Get(logger logger.Logger) Config {
	once.Do(func() {
		logger.Info("Loading config", "file", configFileName)

		bytes, err := os.ReadFile(configFileName)
		if err != nil {
			logger.Error("Failed to read config file", "error", err)
		}

		err = toml.Unmarshal(bytes, &config)
		if err != nil {
			logger.Error("Failed to unmarshal config file", "error", err)
		}

		logger.Info("Config loaded")
	})
	return config
}
