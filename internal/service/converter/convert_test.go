package converter

import (
	"currency-rates/internal/service/converter/mocks"
	"currency-rates/pkg/logger"
	"errors"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	l := logger.New("debug")

	t.Run("Convert success", func(t *testing.T) {
		t.Parallel()

		provider := mocks.NewProvider(t)
		service := New(provider, l)

		provider.On("GetRate", "BTC", "USD").Return(decimal.NewFromInt(1000), nil)

		rate, err := service.Convert(decimal.NewFromInt(1), "BTC", "USD")
		assert.NoError(t, err)
		assert.Equal(t, decimal.NewFromInt(1000), rate)

	})

	t.Run("Convert error", func(t *testing.T) {
		t.Parallel()

		provider := mocks.NewProvider(t)
		service := New(provider, l)

		provider.On("GetRate", "BTC", "USD").Return(decimal.Zero, errors.New("error"))

		rate, err := service.Convert(decimal.NewFromInt(1), "BTC", "USD")
		assert.Error(t, err)
		assert.Equal(t, decimal.Zero, rate)

	})
}
