# Currency Converter Command Line Tool

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/Je33/currency-converter)
[![GitHub Actions](https://img.shields.io/github/actions/workflow/status/Je33/currency-converter/test.yml?style=flat-square)](https://github.com/Je33/currency-converter/actions/workflows/test.yml)


This tool provide the simplest way to get currency rate and convert one currency amount to another 


## Usage

Just run bin file with 3 arguments: amount (float), base currency (string) and target currency (string)

 ```sh
 ./converter 1.25 BTC USD
 ```

And you will get the result:

```sh
1.25 BTC = 10903.75 USD
```

### Providers

#### CoinMarketCap

Converter supports CoinMarketCap API provider
[Learn more](https://pkg.go.dev/github.com/Je33/currency-converter/internal/provider/coinmarketcap).
