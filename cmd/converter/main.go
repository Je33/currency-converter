package main

import (
	"currency-rates/internal/config"
	"currency-rates/internal/provider/coinmarketcap"
	"currency-rates/internal/service/converter"
	"currency-rates/internal/transport/std"
	"currency-rates/pkg/logger"
)

func main() {
	// logger instance
	l := logger.New("debug")

	// get global config
	c := config.Get(l)

	// init provider
	provider := coinmarketcap.New(c.CoinMarketCap, l)

	// init service
	service := converter.New(provider, l)

	// init transport
	transport := std.New(service, l)

	// handle command-lina arguments and convert
	err := transport.Convert()
	if err != nil {
		l.Error("Failed to convert", "error", err)
	}
}
