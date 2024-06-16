package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type GetRateStatus struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
	Notice       string    `json:"notice"`
}

type GetRateDataQuoteCurrency struct {
	Price       decimal.Decimal `json:"price"`
	LastUpdated time.Time       `json:"last_updated"`
}

type GetRateData struct {
	Symbol      string                              `json:"symbol"`
	ID          string                              `json:"id"`
	Name        string                              `json:"name"`
	Amount      decimal.Decimal                     `json:"amount"`
	LastUpdated time.Time                           `json:"last_updated"`
	Quote       map[string]GetRateDataQuoteCurrency `json:"quote"`
}

type GetRateResponse struct {
	Data   map[string]GetRateData `json:"data"`
	Status GetRateStatus          `json:"status"`
}
