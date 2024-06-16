package converter

import (
	"currency-converter/pkg/logger"
	"github.com/shopspring/decimal"
)

//go:generate mockery --dir . --name Provider --output ./mocks --case=underscore
type Provider interface {
	GetRate(base string, target string) (decimal.Decimal, error)
}

type Service struct {
	provider Provider
	logger   logger.Logger
}

func New(provider Provider, logger logger.Logger) *Service {
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
