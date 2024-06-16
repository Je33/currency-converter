package coinmarketcap

import (
	"currency-converter/internal/provider/coinmarketcap/model"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"net/http"
	"strings"
)

func (p *Provider) GetRate(base string, target string) (decimal.Decimal, error) {
	p.logger.Info("CoinMarketCap provider: GetRate", "base", base, "target", target)

	endpoint := "/v2/tools/price-conversion"

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s%s?amount=1&symbol=%s&convert=%s",
			p.config.APIUrl,
			endpoint,
			strings.ToUpper(base),
			strings.ToUpper(target),
		),
		nil,
	)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c")

	p.logger.Debug("CoinMarketCap provider: Sending request", "url", req.URL.String(), "method", req.Method)
	resp, err := p.client.Do(req)
	if err != nil {
		p.logger.Error("CoinMarketCap provider: Failed to send request", "error", err)
		return decimal.Zero, fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			p.logger.Error("CoinMarketCap provider: Failed to close response body", "error", err)
		}
	}()

	p.logger.Debug("CoinMarketCap provider: Received response", "status", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		p.logger.Error("CoinMarketCap provider: Failed to read response body", "error", err)
		return decimal.Zero, fmt.Errorf("failed to read response body: %w", err)
	}
	p.logger.Debug("CoinMarketCap provider: Response body", "body", string(body))

	var respData model.GetRateResponse
	err = json.Unmarshal(body, &respData)
	if err != nil {
		p.logger.Error("CoinMarketCap provider: Failed to unmarshal response body", "error", err)
		return decimal.Zero, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if respData.Status.ErrorCode != 0 || respData.Data == nil {
		p.logger.Error("CoinMarketCap provider: Failed to get rate", "status", respData.Status)
		return decimal.Zero, fmt.Errorf("failed to get rate: %s", respData.Status.ErrorMessage)
	}

	baseCurrency, found := respData.Data[base]
	if !found {
		p.logger.Error("CoinMarketCap provider: Base currency not found in data", "base", base)
		return decimal.Zero, fmt.Errorf("base currency not found in quota: %s", base)
	}

	rate, found := baseCurrency.Quote[target]
	if !found {
		p.logger.Error("CoinMarketCap provider: Target currency not found in quota", "target", target)
		return decimal.Zero, fmt.Errorf("target currency not found in quota: %s", target)
	}

	return rate.Price, nil
}
