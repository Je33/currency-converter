package converter

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func (s *Service) Convert(amount decimal.Decimal, base string, target string) (decimal.Decimal, error) {
	rate, err := s.provider.GetRate(base, target)
	if err != nil {
		s.logger.Error("Rate Service: Failed to get rate", "error", err)
		return decimal.Zero, fmt.Errorf("failed to get rate: %w", err)
	}

	return amount.Mul(rate), nil
}
