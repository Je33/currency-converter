package main

import (
	"currency-converter/internal/config"
	"currency-converter/internal/provider/coinmarketcap"
	"currency-converter/internal/service/converter"
	"currency-converter/internal/transport/std"
	"currency-converter/pkg/logger"
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
