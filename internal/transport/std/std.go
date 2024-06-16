package std

import (
	"currency-rates/pkg/logger"
	"flag"
	"fmt"
	"github.com/shopspring/decimal"
)

type Service interface {
	Convert(amount decimal.Decimal, base string, target string) (decimal.Decimal, error)
}

type Transport struct {
	converter Service
	logger    logger.Logger
}

func New(converter Service, logger logger.Logger) *Transport {
	return &Transport{
		converter: converter,
		logger:    logger,
	}
}

func (t *Transport) Convert() error {
	flag.Parse()

	args := flag.Args()
	if len(args) != 3 {
		return fmt.Errorf("please specify amount, base and target currency, example: 1.25 BTC USD")
	}

	amount, err := decimal.NewFromString(args[0])
	if err != nil {
		return fmt.Errorf("please specify a valid amount, example: 1.25")
	}

	base := args[1]
	target := args[2]

	converted, err := t.converter.Convert(amount, base, target)
	if err != nil {
		t.logger.Error("Failed to convert", "error", err)
		return fmt.Errorf("failed to convert: %w", err)
	}

	fmt.Printf("%s %s = %s %s\n", amount.String(), base, converted.String(), target)

	return nil
}
